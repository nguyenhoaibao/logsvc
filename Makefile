build:
	protoc -I proto/ proto/logsvc.proto --go_out=plugins=gprc:proto
