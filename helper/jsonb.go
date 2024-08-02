package helper

import (
    "encoding/json"
    "umkm/model/domain"
)

func RawMessageToJSONB(raw json.RawMessage) (domain.JSONB, error) {
    var jsonb domain.JSONB
    err := json.Unmarshal(raw, &jsonb)
    if err != nil {
        return nil, err
    }
    return jsonb, nil
}
