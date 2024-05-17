/*
  This file contain the code that will extract data on request.
  It will take access token and get required data.
*/

package services

import (
	"encoding/json"
	"metric-minder-engine/db"
	cErrors "metric-minder-engine/errors"
	httpClient "metric-minder-engine/http"
	"metric-minder-engine/models"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

const BASE_URL = "https://analyticsdata.googleapis.com/v1beta/properties"

type DataExtracter struct {
	statsDb db.StatsDB
}

func NewDataExtracter(db db.StatsDB) DataExtracter {
	return DataExtracter{
		statsDb: db,
	}
}

func (de DataExtracter) GetAndSaveQuickStats(accessToken, propertyID string) error {
	t := time.Now().String()
	tSplit := strings.Split(t, " ")

	if len(tSplit) < 1 {
		return cErrors.NewError("access token invalid")
	}

	today := tSplit[0]

	apiEndpoint := BASE_URL + "/" + propertyID + ":runReport"

	h := map[string]string{
		"Authorization": "Bearer " + accessToken,
		"Accept":        "application/json",
		"Content-Type":  "application/json",
	}

	requestBody := map[string]interface{}{
		"metrics": []map[string]string{
			{"name": "newUsers"},
			{"name": "screenPageViews"},
			{"name": "screenPageViewsPerSession"},
			{"name": "sessions"},
			{"name": "totalUsers"},
		},
		"dateRanges": []map[string]string{
			{
				"startDate": today,
				"endDate":   today,
			},
		},
		"keepEmptyRows": true,
	}

	bytes, err := httpClient.HTTPRequest(http.MethodPost, apiEndpoint, requestBody, nil, h)
	if err != nil {
		log.Error().Msg("Error: " + err.Error())
	}

	resp := models.GAResp{}

	err = json.Unmarshal(bytes, &resp)
	if err != nil {
		log.Error().Msg("Error: " + err.Error())
	}

	stats := models.QuickStats{}

	nV := resp.Rows[0].MetricValues[0].Value
	newVisitors, err := strconv.Atoi(nV)
	if err != nil {
		log.Error().Msg("Error: " + err.Error())
		return err
	}

	stats.NewVisitors = newVisitors

	pV := resp.Rows[0].MetricValues[1].Value
	pageViews, err := strconv.Atoi(pV)
	if err != nil {
		log.Error().Msg("Error: " + err.Error())
		return err
	}

	stats.PageViews = pageViews

	// Pages Per Visit
	ppV := resp.Rows[0].MetricValues[2].Value
	pagesPerVisit, err := strconv.ParseFloat(ppV, 64)
	if err != nil {
		log.Error().Msg("Error: " + err.Error())
		return err
	}

	stats.PagesPerVisit = pagesPerVisit

	tV := resp.Rows[0].MetricValues[3].Value
	totalVisits, err := strconv.Atoi(tV)
	if err != nil {
		log.Error().Msg("Error: " + err.Error())
		return err
	}

	stats.TotalVisits = totalVisits

	v := resp.Rows[0].MetricValues[4].Value
	visitors, err := strconv.Atoi(v)
	if err != nil {
		log.Error().Msg("Error: " + err.Error())
		return err
	}

	stats.Visitors = visitors

	if err := de.statsDb.SaveStats(stats); err != nil {
		log.Error().Msg("Error: " + err.Error())
		return err
	}

	return nil
}
