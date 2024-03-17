generate_proto:
	python -m grpc_tools.protoc -I./protos --python_out=./process_pdf_service --pyi_out=./process_pdf_service --grpc_python_out=./process_pdf_service ./protos/process_pdf_service.proto \
	protoc --go_out=./api --go_opt=paths=source_relative --go-grpc_out=./api --go-grpc_opt=paths=source_relative ./protos/process_pdf_service.proto