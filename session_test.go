package gwda

import (
	"encoding/base64"
	"strings"
	"testing"
)

func TestSession_GetActiveSession(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	sessionInfo, err := s.GetActiveSession()
	if err != nil {
		t.Fatal(err)
	}
	if sessionInfo.SessionID == "" {
		t.Fatal("session id should not be empty")
	}

	t.Logf("\n当前 App 的 名称:\t\t%s\n"+
		"当前 App 的 BundleId:\t%s\n"+
		"当前设备的系统版本:\t\t%s\n",
		sessionInfo.Capabilities.BrowserName, sessionInfo.Capabilities.CFBundleIdentifier, sessionInfo.Capabilities.SdkVersion)
	// t.Log("当前 App 的 BundleId:", sessionInfo.Capabilities.CFBundleIdentifier)
}

func TestSession_DeleteSession(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	err = s.Delete()
	if err != nil {
		t.Fatal(err)
	}

	// t.Log(c.ActiveAppInfo())
	err = s.SiriActivate("打开 微信")
	if err == nil {
		t.Fatal("It should not be nil")
	}
	if !strings.EqualFold(err.Error(), "invalid session id: Session does not exist") {
		t.Fatal(err)
	}

	// t.Log(s.SiriOpenURL("weixin://"))
}

func TestSession_DeviceInfo(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	s, err := c.NewSession("com.apple.Preferences")
	if err != nil {
		t.Fatal(err)
	}
	dInfo, err := s.DeviceInfo()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(dInfo)
	t.Log(dInfo.Name)
	t.Log(dInfo.UserInterfaceStyle)
}

func TestSession_BatteryInfo(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	s, err := c.NewSession("com.apple.Preferences")
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	batteryInfo, err := s.BatteryInfo()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(batteryInfo)
	t.Log(batteryInfo.Level)
	t.Log(batteryInfo.State)
}

func TestSession_WindowSize(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	s, err := c.NewSession("com.apple.Preferences")
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	windowSize, err := s.WindowSize()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(windowSize)
}

func TestSession_Screen(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	s, err := c.NewSession("com.apple.Preferences")
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	screen, err := s.Screen()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(screen)
}

func TestSession_Scale(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	s, err := c.NewSession("com.apple.Preferences")
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	scale, err := s.Scale()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(scale)
}

func TestSession_StatusBarSize(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	s, err := c.NewSession("com.apple.Preferences")
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	statusBarSize, err := s.StatusBarSize()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(statusBarSize)
	t.Log(statusBarSize.Width)
	t.Log(statusBarSize.Height)
}

func TestSession_ActiveAppInfo(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	appInfo, err := s.ActiveAppInfo()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(appInfo)
}

func TestSession_ActiveAppsList(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	appsList, err := s.ActiveAppsList()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(appsList)
}

func TestSession_Tap(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	err = s.Tap(230, 130)
	if err != nil {
		t.Fatal(err)
	}
	err = s.Tap(210, 290)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSession_DoubleTap(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	err = s.DoubleTap(230, 130)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSession_TouchAndHold(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	// err = s.TouchAndHold(210, 290)
	err = s.TouchAndHold(230, 130)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSession_AppLaunch(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	// bundleId = "com.apple.AppStore"
	_ = s.AppTerminate(bundleId)
	// bool 类型初始值为 false，在设置启动操作选项需要主动设置 ShouldWaitForQuiescence 为 true （如果需要的话）
	launchOpt := WDAAppLaunchOption{
		ShouldWaitForQuiescence: true,
		Arguments:               []string{"-AppleLanguages", "(Russian)"},
		// Arguments:               []string{"-AppleLanguages", "(ru)"},
		// Environment:             map[string]string{"DYLD_PRINT_STATISTICS": "YES"},
	}
	_ = launchOpt
	err = s.AppLaunch(bundleId, launchOpt)
	// err = s.AppLaunch(bundleId)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSession_AppTerminate(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	bundleId := "com.apple.Preferences"
	s, err := c.NewSession(bundleId)
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	err = s.AppTerminate(bundleId)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSession_AppState(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	// bundleId := "com.apple.Preferences"
	s, err := c.NewSession(bundleId)
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	state, err := s.AppState(bundleId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(state)
	t.Log("app 是否前台活动中", state == WDAAppRunningFront)
}

func TestSession_SendKeys(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	bundleId := "com.apple.Preferences"
	s, err := c.NewSession(bundleId)
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	err = s.SendKeys(bundleId)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSession_FindElement(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	bundleId := "com.apple.Preferences"
	s, err := c.NewSession(bundleId)
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	// elements, err := s.FindElements("partial link text", "label=看一看")
	elements, err := s.FindElements("partial link text", "label=发现")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(elements)

	// if len(elements) == 1 {
	// 	err := elements[0].Click()
	// 	t.Log(err)
	// }
}

func TestSession_Locked(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	locked, err := s.IsLocked()
	if err != nil {
		t.Fatal(err)
	}
	if locked {
		t.Log("设备处于 屏幕锁定 界面")
	} else {
		t.Log("设备已屏幕解锁")
	}
}

func TestSession_Unlock(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	bundleId := "com.apple.Preferences"
	s, err := c.NewSession(bundleId)
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	err = s.Unlock()
	if err != nil {
		t.Fatal(err)
	}
}

func TestSession_Lock(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	bundleId := "com.apple.Preferences"
	s, err := c.NewSession(bundleId)
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	err = s.Lock()
	if err != nil {
		t.Fatal(err)
	}
}

func TestSession_Activate(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	// bundleId := "com.apple.Preferences"
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	err = s.AppActivate(bundleId)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSession_DeactivateApp(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	bundleId := "com.apple.Preferences"
	s, err := c.NewSession(bundleId)
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	err = s.DeactivateApp(20.1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSession_SetPasteboardForPlaintext(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	bundleId := "com.apple.Preferences"
	s, err := c.NewSession(bundleId)
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	err = s.SetPasteboardForPlaintext("abcd1234")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSession_SetPasteboardForImage(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	bundleId := "com.apple.Preferences"
	s, err := c.NewSession(bundleId)
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	err = s.SetPasteboardForImage("/Users/hero/Documents/leixipaopao/IMG_5246.JPG")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSession_SetPasteboardForUrl(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	bundleId := "com.apple.Preferences"
	s, err := c.NewSession(bundleId)
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	err = s.SetPasteboardForUrl("http://baidu.com")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSession_SetPasteboard(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	bundleId := "com.apple.Preferences"
	s, err := c.NewSession(bundleId)
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	encodedContent := base64.URLEncoding.EncodeToString([]byte("https://www.google.com"))
	err = s.SetPasteboard(WDAContentTypeUrl, encodedContent)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSession_PressHomeButton(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	bundleId := "com.apple.Preferences"
	s, err := c.NewSession(bundleId)
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	err = s.PressHomeButton()
	if err != nil {
		t.Fatal(err)
	}
}

func TestSession_PressVolumeUpButton(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	bundleId := "com.apple.Preferences"
	s, err := c.NewSession(bundleId)
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	err = s.PressVolumeUpButton()
	if err != nil {
		t.Fatal(err)
	}
}

func TestSession_PressVolumeDownButton(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	bundleId := "com.apple.Preferences"
	s, err := c.NewSession(bundleId)
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	err = s.PressVolumeDownButton()
	if err != nil {
		t.Fatal(err)
	}
}

func TestSession_SiriActivate(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	err = s.SiriActivate("查一下附近的餐厅")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSession_SiriOpenURL(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	err = s.SiriOpenURL("weixin://")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSession_Source(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	bundleId := "com.apple.Preferences"
	s, err := c.NewSession(bundleId)
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	// sTree, err := s.Source()
	sTree, err := s.Source(true)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sTree)
}

func TestSession_AccessibleSource(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	bundleId := "com.apple.Preferences"
	s, err := c.NewSession(bundleId)
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	source, err := s.AccessibleSource()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(source)
}

func TestTmpSession(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	s.tttTmp()
	// _ = s
}
