.PHONY: update build tidy clone-wiki clone-wiki-only update-wiki

tidy:
	go mod tidy

update: tidy
	go build -o gfp cmd/main.go && ./gfp

build: tidy
	go build -v -o gfp cmd/main.go

# clone-wiki-only: 仅克隆 wiki 仓库（用于轻量复测时读取当前 lists）
clone-wiki-only:
	rm -rf ../wiki || true
	git clone https://github.com/feitianyul/free-proxy-list/wiki.git ../wiki || true
	mkdir -p ../wiki/lists

# clone-wiki: 克隆 wiki 并将本地 list/* 拷贝到 wiki/lists（全量更新后用）
clone-wiki: clone-wiki-only
	cp -r list/* ../wiki/lists/ 2>/dev/null || true


# update-wiki: build Home.md and push to wiki (assumes wiki/ already contains lists/)
update-wiki:
	./update_wiki.sh
