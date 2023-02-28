/*
 * Produced: Mon Feb 27 2023
 * Author: Alec M.
 * GitHub: https://amattu.com/links/github
 * Copyright: (C) 2023 Alec M.
 * License: License GNU Affero General Public License v3.0
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package constants

const (
	BMW         = "BMW"
	MINI        = "MINI"
	ROLLS_ROYCE = "ROLLS-ROYCE"

	USER_AGENT   = "Dart/2.16 (dart:io)"
	X_USER_AGENT = "android(SP1A.210812.016.C1);{%s};{%s};{%s}"

	NORTH_AMERICA = "na"
	CHINA         = "cn"
	REST_OF_WORLD = "row"

	SERVICE_CHARGING_STATISTICS_URL = "CHARGING_STATISTICS"
	SERVICE_CHARGING_SESSIONS_URL   = "CHARGING_SESSIONS"
	SERVICE_CHARGING_PROFILE        = "CHARGING_PROFILE"

	ATTR_STATE             = "state"
	ATTR_CAPABILITIES      = "capabilities"
	ATTR_ATTRIBUTES        = "attributes"
	ATTR_CHARGING_SETTINGS = "charging_settings"
)

const (
	AUTH_CHINA_PUBLIC_KEY_URL = "/eadrax-coas/v1/cop/publickey"
	AUTH_CHINA_LOGIN_URL      = "/eadrax-coas/v2/login/pwd"
	AUTH_CHINA_TOKEN_URL      = "/eadrax-coas/v1/oauth/token"

	OAUTH_CONFIG_URL = "/eadrax-ucs/v1/presentation/oauth/config"

	VEHICLES_URL      = "/eadrax-vcs/v4/vehicles"
	VEHICLE_STATE_URL = VEHICLES_URL + "/state"

	REMOTE_SERVICE_BASE_URL     = "/eadrax-vrccs/v3/presentation/remote-commands"
	REMOTE_SERVICE_URL          = REMOTE_SERVICE_BASE_URL + "/{vin}/%s"
	REMOTE_SERVICE_STATUS_URL   = REMOTE_SERVICE_BASE_URL + "/eventStatus?eventId=%s"
	REMOTE_SERVICE_POSITION_URL = REMOTE_SERVICE_BASE_URL + "/eventPosition?eventId=%s"

	VEHICLE_CHARGING_DETAILS_URL      = "/eadrax-crccs/v2/vehicles"
	VEHICLE_CHARGING_BASE_URL         = "/eadrax-crccs/v1/vehicles/%s"
	VEHICLE_CHARGING_SETTINGS_SET_URL = VEHICLE_CHARGING_BASE_URL + "/charging-settings"
	VEHICLE_CHARGING_PROFILE_SET_URL  = VEHICLE_CHARGING_BASE_URL + "/charging-profile"
	VEHICLE_CHARGING_START_STOP_URL   = VEHICLE_CHARGING_BASE_URL + "/%s"
	VEHICLE_CHARGING_STATISTICS_URL   = "/eadrax-chs/v1/charging-statistics"
	VEHICLE_CHARGING_SESSIONS_URL     = "/eadrax-chs/v1/charging-sessions"

	VEHICLE_IMAGE_URL = "/eadrax-ics/v3/presentation/vehicles/{vin}/images?carView=%s"
	VEHICLE_POI_URL   = "/eadrax-dcs/v1/send-to-car/send-to-car"
)

var VCS_SERVER_URLS = map[string]string{
	NORTH_AMERICA: "cocoapi.bmwgroup.us",
	REST_OF_WORLD: "cocoapi.bmwgroup.com",
	CHINA:         "myprofile.bmw.com.cn",
}

var OCP_APIM_KEYS = map[string]string{
	NORTH_AMERICA: "MzFlMTAyZjUtNmY3ZS03ZWYzLTkwNDQtZGRjZTYzODkxMzYy",
	REST_OF_WORLD: "NGYxYzg1YTMtNzU4Zi1hMzdkLWJiYjYtZjg3MDQ0OTRhY2Zh",
}

var AES_KEYS = map[string]map[string]string{
	CHINA: {
		"key": "UzJUdzEwdlExWGYySmxLYQ==",
		"iv":  "dTFGUDd4ZWRrQWhMR3ozVQ==",
	},
}

var APP_VERSIONS = map[string]string{
	NORTH_AMERICA: "2.12.0(19883)",
	REST_OF_WORLD: "2.12.0(19883)",
	CHINA:         "2.3.0(13603)",
}

var GARAGE_SERVER_URLS = map[string]map[string]string{
	BMW: {
		NORTH_AMERICA: "mygarage.bmwusa.com",
	},
}
