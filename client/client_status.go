package client

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

type ClientStatusService struct {
	client        *CommvaultClient
	isActive      bool
	isActiveLock  sync.Mutex
	checkPeriod   time.Duration
	ticker        *time.Ticker
	lastCheckTime time.Time
}

func NewClientStatusService(client *CommvaultClient, period time.Duration) *ClientStatusService {
	return &ClientStatusService{
		client:       client,
		checkPeriod:  period,
		isActiveLock: sync.Mutex{},
		isActive:     false,
		ticker:       time.NewTicker(period),
	}
}

func (c *ClientStatusService) GetIsActive() bool {
	c.isActiveLock.Lock()
	defer c.isActiveLock.Unlock()
	return c.isActive
}

func (c *ClientStatusService) GetStatus() ClientStatus {
	c.isActiveLock.Lock()
	defer c.isActiveLock.Unlock()
	return ClientStatus{
		IsActive:        c.GetIsActive(),
		LastStatusCheck: c.lastCheckTime,
	}
}

func (c *ClientStatusService) SetStatus(active bool) {
	c.isActiveLock.Lock()
	defer c.isActiveLock.Unlock()
	c.isActive = active
}

func (c *ClientStatusService) Start(ctx context.Context) {
	go c.executor(ctx)
}

func (c *ClientStatusService) Stop(ctx context.Context) {
	c.ticker.Stop()
}

func (c *ClientStatusService) executor(ctx context.Context) {
	for {
		select {
		case <-c.ticker.C:
			err := c.CheckStatus()
			if err != nil {
				fmt.Printf("[error]: failed to check API endpoint status\n%+v\n", err)
				continue
			}
			var s string
			if c.GetIsActive() {
				s = "Active"
			} else {
				s = "Standby"
			}
			fmt.Printf("[info] Commvault API status: %s\n", s)
		case <-ctx.Done():
			return
		}
	}
}

func (c *ClientStatusService) CheckStatus() error {
	baseUrl, err := c.client.Url(nil, nil)
	if err != nil {
		return err
	}

	var port int
	if port, err = strconv.Atoi(baseUrl.Port()); err != nil || port == 0 {
		if baseUrl.Scheme == "http" {
			port = 80
		} else if baseUrl.Scheme == "https" {
			port = 443
		} else {
			return fmt.Errorf("invalid port number")
		}
	}

	fmt.Printf("[debug] Status check: %s:%d\n", baseUrl.Hostname(), port)
	var previousStatus bool
	// conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", baseUrl.Hostname(), port))
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", baseUrl.Hostname(), port), 3*time.Second)
	previousStatus = c.isActive
	if err != nil {
		c.SetStatus(false)
		fmt.Printf("[debug] Status check: failed\n")
	} else {
		c.SetStatus(true)
		fmt.Printf("[debug] Status check: successfull\n")
	}

	if previousStatus && !c.isActive {
		fmt.Println("[debug] Trigger token flush")
		go func() {
			c.client.tokenFlushTrigger <- true
		}()

	} else if !previousStatus && c.isActive {
		fmt.Println("[debug] Trigger token renewal")
		go func() {
			c.client.tokenManualTrigger <- true
		}()
	}

	defer func() {
		if conn != nil {
			err := conn.Close()
			if err != nil {
				fmt.Printf("failed to close connection: %s\n", err)
			}
		}
	}()

	return nil
}
