package infoapi

import (
	"net/http"

	"osAnalytics/internal/getcpuinfo"
	"osAnalytics/internal/getgeneralinfo"
	"osAnalytics/internal/getmeminfo"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", mainPageHandler)
	r.GET("/system", getGeneralInfoHandler)
	r.GET("/processor", getCPUInfoHandler)
	r.GET("/memory", getMemInfoHandler)
}

func mainPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Main page",
	})
}

func getCPUInfoHandler(c *gin.Context) {
	var cpuInfo getcpuinfo.CPUInfo = getcpuinfo.CPUInfo{}
	err := cpuInfo.GetInfo()

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.HTML(http.StatusOK, "cpu.tmpl", gin.H{
		"title":   "CPU information",
		"cpuInfo": cpuInfo,
	})
}

func getGeneralInfoHandler(c *gin.Context) {
	var sysInfo getgeneralinfo.SystemInfo = getgeneralinfo.SystemInfo{}
	err := sysInfo.GetInfo()

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.HTML(http.StatusOK, "system.tmpl", gin.H{
		"title":      "System information",
		"systemInfo": sysInfo,
	})
}

func getMemInfoHandler(c *gin.Context) {
	var memInfo getmeminfo.MemInfo = getmeminfo.MemInfo{}
	err := memInfo.GetInfo()

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.HTML(http.StatusOK, "mem.tmpl", gin.H{
		"title":   "System information",
		"memInfo": memInfo,
	})
}
