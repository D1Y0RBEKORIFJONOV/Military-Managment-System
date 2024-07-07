gen-soldiers:
	protoc -I proto proto/soldiers/soldiers.proto --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative
gen-storehouses:
	protoc -I proto proto/storehouses/storehouses.proto --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go --go-grpc_opt=paths=source_relative
gen-teachers:
	protoc -I proto proto/teacher/teacher.proto --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go --go-grpc_opt=paths=source_relative
gen-order:
	protoc -I proto proto/order/order.proto --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go --go-grpc_opt=paths=source_relative