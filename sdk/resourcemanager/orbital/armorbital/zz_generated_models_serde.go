//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armorbital

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type AvailableContactsListResult.
func (a AvailableContactsListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", a.NextLink)
	populate(objectMap, "value", a.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type AvailableGroundStationListResult.
func (a AvailableGroundStationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", a.NextLink)
	populate(objectMap, "value", a.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type CloudErrorBody.
func (c CloudErrorBody) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", c.Code)
	populate(objectMap, "details", c.Details)
	populate(objectMap, "message", c.Message)
	populate(objectMap, "target", c.Target)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ContactInstanceProperties.
func (c ContactInstanceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "endAzimuthDegrees", c.EndAzimuthDegrees)
	populate(objectMap, "endElevationDegrees", c.EndElevationDegrees)
	populate(objectMap, "maximumElevationDegrees", c.MaximumElevationDegrees)
	populateTimeRFC3339(objectMap, "rxEndTime", c.RxEndTime)
	populateTimeRFC3339(objectMap, "rxStartTime", c.RxStartTime)
	populate(objectMap, "startAzimuthDegrees", c.StartAzimuthDegrees)
	populate(objectMap, "startElevationDegrees", c.StartElevationDegrees)
	populateTimeRFC3339(objectMap, "txEndTime", c.TxEndTime)
	populateTimeRFC3339(objectMap, "txStartTime", c.TxStartTime)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ContactInstanceProperties.
func (c *ContactInstanceProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "endAzimuthDegrees":
			err = unpopulate(val, &c.EndAzimuthDegrees)
			delete(rawMsg, key)
		case "endElevationDegrees":
			err = unpopulate(val, &c.EndElevationDegrees)
			delete(rawMsg, key)
		case "maximumElevationDegrees":
			err = unpopulate(val, &c.MaximumElevationDegrees)
			delete(rawMsg, key)
		case "rxEndTime":
			err = unpopulateTimeRFC3339(val, &c.RxEndTime)
			delete(rawMsg, key)
		case "rxStartTime":
			err = unpopulateTimeRFC3339(val, &c.RxStartTime)
			delete(rawMsg, key)
		case "startAzimuthDegrees":
			err = unpopulate(val, &c.StartAzimuthDegrees)
			delete(rawMsg, key)
		case "startElevationDegrees":
			err = unpopulate(val, &c.StartElevationDegrees)
			delete(rawMsg, key)
		case "txEndTime":
			err = unpopulateTimeRFC3339(val, &c.TxEndTime)
			delete(rawMsg, key)
		case "txStartTime":
			err = unpopulateTimeRFC3339(val, &c.TxStartTime)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ContactListResult.
func (c ContactListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", c.NextLink)
	populate(objectMap, "value", c.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ContactParameters.
func (c ContactParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "contactProfile", c.ContactProfile)
	populateTimeRFC3339(objectMap, "endTime", c.EndTime)
	populate(objectMap, "groundStationName", c.GroundStationName)
	populateTimeRFC3339(objectMap, "startTime", c.StartTime)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ContactParameters.
func (c *ContactParameters) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "contactProfile":
			err = unpopulate(val, &c.ContactProfile)
			delete(rawMsg, key)
		case "endTime":
			err = unpopulateTimeRFC3339(val, &c.EndTime)
			delete(rawMsg, key)
		case "groundStationName":
			err = unpopulate(val, &c.GroundStationName)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateTimeRFC3339(val, &c.StartTime)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ContactProfile.
func (c ContactProfile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", c.Etag)
	populate(objectMap, "id", c.ID)
	populate(objectMap, "location", c.Location)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "properties", c.Properties)
	populate(objectMap, "systemData", c.SystemData)
	populate(objectMap, "tags", c.Tags)
	populate(objectMap, "type", c.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ContactProfileLink.
func (c ContactProfileLink) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "channels", c.Channels)
	populate(objectMap, "direction", c.Direction)
	populate(objectMap, "eirpdBW", c.EirpdBW)
	populate(objectMap, "gainOverTemperature", c.GainOverTemperature)
	populate(objectMap, "polarization", c.Polarization)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ContactProfileListResult.
func (c ContactProfileListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", c.NextLink)
	populate(objectMap, "value", c.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ContactProfilesProperties.
func (c ContactProfilesProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "autoTrackingConfiguration", c.AutoTrackingConfiguration)
	populate(objectMap, "eventHubUri", c.EventHubURI)
	populate(objectMap, "links", c.Links)
	populate(objectMap, "minimumElevationDegrees", c.MinimumElevationDegrees)
	populate(objectMap, "minimumViableContactDuration", c.MinimumViableContactDuration)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ContactsProperties.
func (c ContactsProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "contactProfile", c.ContactProfile)
	populate(objectMap, "endAzimuthDegrees", c.EndAzimuthDegrees)
	populate(objectMap, "endElevationDegrees", c.EndElevationDegrees)
	populate(objectMap, "errorMessage", c.ErrorMessage)
	populate(objectMap, "groundStationName", c.GroundStationName)
	populate(objectMap, "maximumElevationDegrees", c.MaximumElevationDegrees)
	populateTimeRFC3339(objectMap, "reservationEndTime", c.ReservationEndTime)
	populateTimeRFC3339(objectMap, "reservationStartTime", c.ReservationStartTime)
	populateTimeRFC3339(objectMap, "rxEndTime", c.RxEndTime)
	populateTimeRFC3339(objectMap, "rxStartTime", c.RxStartTime)
	populate(objectMap, "startAzimuthDegrees", c.StartAzimuthDegrees)
	populate(objectMap, "startElevationDegrees", c.StartElevationDegrees)
	populate(objectMap, "status", c.Status)
	populateTimeRFC3339(objectMap, "txEndTime", c.TxEndTime)
	populateTimeRFC3339(objectMap, "txStartTime", c.TxStartTime)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ContactsProperties.
func (c *ContactsProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "contactProfile":
			err = unpopulate(val, &c.ContactProfile)
			delete(rawMsg, key)
		case "endAzimuthDegrees":
			err = unpopulate(val, &c.EndAzimuthDegrees)
			delete(rawMsg, key)
		case "endElevationDegrees":
			err = unpopulate(val, &c.EndElevationDegrees)
			delete(rawMsg, key)
		case "errorMessage":
			err = unpopulate(val, &c.ErrorMessage)
			delete(rawMsg, key)
		case "groundStationName":
			err = unpopulate(val, &c.GroundStationName)
			delete(rawMsg, key)
		case "maximumElevationDegrees":
			err = unpopulate(val, &c.MaximumElevationDegrees)
			delete(rawMsg, key)
		case "reservationEndTime":
			err = unpopulateTimeRFC3339(val, &c.ReservationEndTime)
			delete(rawMsg, key)
		case "reservationStartTime":
			err = unpopulateTimeRFC3339(val, &c.ReservationStartTime)
			delete(rawMsg, key)
		case "rxEndTime":
			err = unpopulateTimeRFC3339(val, &c.RxEndTime)
			delete(rawMsg, key)
		case "rxStartTime":
			err = unpopulateTimeRFC3339(val, &c.RxStartTime)
			delete(rawMsg, key)
		case "startAzimuthDegrees":
			err = unpopulate(val, &c.StartAzimuthDegrees)
			delete(rawMsg, key)
		case "startElevationDegrees":
			err = unpopulate(val, &c.StartElevationDegrees)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, &c.Status)
			delete(rawMsg, key)
		case "txEndTime":
			err = unpopulateTimeRFC3339(val, &c.TxEndTime)
			delete(rawMsg, key)
		case "txStartTime":
			err = unpopulateTimeRFC3339(val, &c.TxStartTime)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OperationListResult.
func (o OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceIDListResult.
func (r ResourceIDListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", r.NextLink)
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Spacecraft.
func (s Spacecraft) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", s.Etag)
	populate(objectMap, "id", s.ID)
	populate(objectMap, "location", s.Location)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "systemData", s.SystemData)
	populate(objectMap, "tags", s.Tags)
	populate(objectMap, "type", s.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SpacecraftListResult.
func (s SpacecraftListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SpacecraftsProperties.
func (s SpacecraftsProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "authorizationStatus", s.AuthorizationStatus)
	populate(objectMap, "authorizationStatusExtended", s.AuthorizationStatusExtended)
	populate(objectMap, "links", s.Links)
	populate(objectMap, "noradId", s.NoradID)
	populate(objectMap, "titleLine", s.TitleLine)
	populate(objectMap, "tleLine1", s.TleLine1)
	populate(objectMap, "tleLine2", s.TleLine2)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SystemData.
func (s SystemData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemData.
func (s *SystemData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateTimeRFC3339(val, &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type TagsObject.
func (t TagsObject) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "tags", t.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type TrackedResource.
func (t TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", t.ID)
	populate(objectMap, "location", t.Location)
	populate(objectMap, "name", t.Name)
	populate(objectMap, "systemData", t.SystemData)
	populate(objectMap, "tags", t.Tags)
	populate(objectMap, "type", t.Type)
	return json.Marshal(objectMap)
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}