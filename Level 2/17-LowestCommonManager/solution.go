package main

type OrgChart struct {
	Name string
	DirectReports []*OrgChart
}

type OrgInfo struct {
	lowestCommonManager *OrgChart
	numImportantReports int
}

func GetLowestCommonManager(org, reportOne, reportTwo *OrgChart) *OrgChart {
	return getOrgInfo(org, reportOne, reportTwo).lowestCommonManager
}

func getOrgInfo(manager, reportOne, reportTwo *OrgChart) OrgInfo {
	numImportantReports := 0
	for _, directReport := range manager.DirectReports {
		orgInfo := getOrgInfo(directReport, reportOne, reportTwo)
		if orgInfo.lowestCommonManager != nil {
			return orgInfo
		}
		numImportantReports += orgInfo.numImportantReports
	}
	if manager == reportOne || manager == reportTwo {
		numImportantReports++
	}
	var lowestCommonManager *OrgChart
	if numImportantReports == 2 {
		lowestCommonManager = manager
	}
	return OrgInfo{
		lowestCommonManager: lowestCommonManager,
		numImportantReports: numImportantReports,
	}
}