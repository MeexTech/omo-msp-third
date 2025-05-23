.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/third/common.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/third/partner.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/third/channel.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/third/motion.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/third/carousel.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/third/recommend.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/third/schedule.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/third/holiday.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/third/netflow.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/third/honor.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/third/app.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/third/kms.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/third/link.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/third/reserve.proto
