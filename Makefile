# https://stackoverflow.com/a/12959694
rwildcard=$(wildcard $1$2) $(foreach d,$(wildcard $1*),$(call rwildcard,$d/,$2))

RPC_PROTOS := $(call rwildcard,rpc/,*.proto)
LIB_PROTOS := $(call rwildcard,util/,*.proto)

RPC_PBGENS := $(RPC_PROTOS:.proto=.twirp.go)
LIB_PBGENS := $(LIB_PROTOS:.proto=.pb.go)

.PRECIOUS: $(RPC_PBGENS) $(LIB_PBGENS)

# 参数 Mfoo.proto=bar/foo 表示 foo.proto 生成的 go 文件所对应的包名是 bar/foo。
#
# 如是在 proto 中引用了其他 proto，生成的 go 文件需要导入对应的包。
# 但 protoc 和 proto-gen-go 无法单独从 proto 文件获取当前项目的包名，
# 最好的办法就是通过 go_package 手工指定，但这样写起来太丑了，所以改用 M 参数。
#
# 如果你自己写了包供别人导入使用，则一定要在 proto 中设置 go_package 选项。
#
# 更多讨论请参考
# https://github.com/golang/protobuf/issues/1158#issuecomment-650694184
#
# $(...) 中的神奇代码是为实现以下替换
# pkg/kv/taishan/taishan.proto => pkg/taishan
%.pb.go: %.proto
	protoc --go_out=M$<=$(patsubst %/,%,$(dir $<)):. $<

# $(...) 中的神奇代码是为实现以下替换
# rpc/pkg/v0/kv.proto => rpc/pkg/v0;util_v0
%.twirp.go: %.proto
	$(eval m=$<=$(join \
			$(patsubst %/,%\;,\
				$(dir $<)\
			),\
			$(subst /v,_v,\
				$(patsubst rpc/%,%,\
					$(patsubst %/,%,$(dir $<))\
				)\
			)\
		))
	protoc --twirp_out=root_package=moocss.com/gaea,validate_enable=true,M$m:. \
		--go_out=M$m:. \
		$<

default: rpc util
	go build -trimpath -mod=readonly

.PHONY: rpc
rpc: $(RPC_PBGENS)
	@exit

.PHONY: util
util: $(LIB_PBGENS)
	@exit

.PHONY: cmd
cmd:
	go install ./cmd/protoc-gen-twirp

.PHONY: clean
clean:
	git clean -x -f -d

.PHONY: rename
rename:
	go run cmd/sniper/main.go rename  --package $(name)

.PHONY: run-public
run-public:
	export APP_ID=GaeaApi; export DEPLOY_ENV=uat; go run main.go server --port=8080;

.PHONY: run-private
run-private:
	export APP_ID=GaeaInternalApi; go run main.go server --port=8080 --internal;

.PHONY: run-job
run-job:
	export APP_ID=GaeaJob; go run main.go job --port=8081;

.PHONY: wire
wire:
	wire ./cmd/server

.PHONY: ent
ent:
	cd internal/data/ && ent generate ./ent/schema	
