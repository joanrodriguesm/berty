// Code generated by berty.tech/core/.scripts/generate-logger.sh

package models

import "go.uber.org/zap"

func logger() *zap.Logger {
	return zap.L().Named("vendor.graphql.models")
}
