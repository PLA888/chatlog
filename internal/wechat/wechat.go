package wechat

import (
	"context"
	"fmt"
	"os"

	"github.com/sjzar/chatlog/internal/wechat/decrypt"
	"github.com/sjzar/chatlog/internal/wechat/key"
	"github.com/sjzar/chatlog/internal/wechat/model"
)

// Account 表示一个微信账号
type Account struct {
	Name        string
	Platform    string
	Version     int
	FullVersion string
	DataDir     string
	Key         string
	PID         uint32
	ExePath     string
	Status      string
}

// NewAccount 创建新的账号对象
func NewAccount(proc *model.Process) *Account {
	return &Account{
		Name:        proc.AccountName,
		Platform:    proc.Platform,
		Version:     proc.Version,
		FullVersion: proc.FullVersion,
		DataDir:     proc.DataDir,
		PID:         proc.PID,
		ExePath:     proc.ExePath,
		Status:      proc.Status,
	}
}

// RefreshStatus 刷新账号的进程状态
func (a *Account) RefreshStatus() error {
	// 查找所有微信进程
	Load()

	process, err := GetProcess(a.Name)
	if err != nil {
		a.Status = model.StatusOffline
		return nil
	}

	if process.AccountName == a.Name {
		// 更新进程信息
		a.PID = process.PID
		a.ExePath = process.ExePath
		a.Platform = process.Platform
		a.Version = process.Version
		a.FullVersion = process.FullVersion
		a.Status = process.Status
		a.DataDir = process.DataDir
	}

	return nil
}

// GetKey 获取账号的密钥
func (a *Account) GetKey(ctx context.Context) (string, error) {
	// 如果已经有密钥，直接返回
	if a.Key != "" {
		return a.Key, nil
	}

	// 刷新进程状态
	if err := a.RefreshStatus(); err != nil {
		return "", fmt.Errorf("failed to refresh process status: %w", err)
	}

	// 检查账号状态
	if a.Status != model.StatusOnline {
		return "", fmt.Errorf("account %s is not online", a.Name)
	}

	// 创建密钥提取器 - 使用新的接口，传入平台和版本信息
	extractor, err := key.NewExtractor(a.Platform, a.Version)
	if err != nil {
		return "", fmt.Errorf("failed to create key extractor: %w", err)
	}

	process, err := GetProcess(a.Name)
	if err != nil {
		return "", fmt.Errorf("failed to get process: %w", err)
	}

	validator, err := decrypt.NewValidator(process.DataDir, process.Platform, process.Version)
	if err != nil {
		return "", fmt.Errorf("failed to create validator: %w", err)
	}

	extractor.SetValidate(validator)

	// 提取密钥
	key, err := extractor.Extract(ctx, process)
	if err != nil {
		return "", err
	}

	// 保存密钥
	a.Key = key
	return key, nil
}

// DecryptDatabase 解密数据库
func (a *Account) DecryptDatabase(ctx context.Context, dbPath, outputPath string) error {
	// 获取密钥
	hexKey, err := a.GetKey(ctx)
	if err != nil {
		return err
	}

	// 创建解密器 - 传入平台信息和版本
	decryptor, err := decrypt.NewDecryptor(a.Platform, a.Version)
	if err != nil {
		return err
	}

	// 创建输出文件
	output, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer output.Close()

	// 解密数据库
	return decryptor.Decrypt(ctx, dbPath, hexKey, output)
}
