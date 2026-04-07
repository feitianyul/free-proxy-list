**中文** | [English (README_EN.md)](README_EN.md)

---

<p align="center">
  <img src="https://img.shields.io/badge/每1小时更新-通过-success">  
  <br>
  <img src="https://img.shields.io/website/https/getfreeproxy.com.svg">
  <img src="https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/total.svg">
  <img src="https://img.shields.io/github/last-commit/feitianyul/free-proxy-list.svg">
  <img src="https://img.shields.io/github/license/feitianyul/free-proxy-list.svg">
  
  <br>
  <br>
  <a href="https://getfreeproxy.com/lists/" title="可用代理列表">可用代理列表</a> | <a href="https://getfreeproxy.com/tools/proxy-checker" title="在线代理检测">免费代理检测</a> | <a href="https://getfreeproxy.com/tools/proxy-protocol-parser" title="代理协议解析">通用代理协议解析</a> | <a href="https://developer.getfreeproxy.com/" title="代理 API">免费代理 API</a>
  <br>
</p>

# 🌎 GetFreeProxy (GFP)：免费代理列表

**GetFreeProxy (GFP)** 是一个开源项目，自动从互联网聚合并校验免费代理，旨在为开发者、研究人员及需要代理服务的用户提供新鲜、可靠、可用的公共代理列表。

列表按小时更新，确保您始终能获取到最新的可用代理。

---

## 📖 项目说明

本项目为开源免费代理聚合与校验工具，从互联网公开源拉取代理并**仅保留 HTTP、HTTPS** 两种类型，经校验后生成列表，供开发者、研究人员等使用。

### 本仓库特点

- **仅保留两种代理**：HTTP、HTTPS，不收录 SOCKS、VMess、Trojan、VLESS、SS/SSR、Hysteria 等其它协议。
- **校验规则**：五域名中**任意 3 个**在 2 秒内成功（HTTP 200）即视为该协议通过。对每条代理访问以下五个地址验证（优先 HEAD，不支持则回退 GET）：
  - `https://www.eastmoney.com/`
  - `https://www.sse.com.cn/`
  - `https://finance.sina.com.cn/`（新浪财经）
  - `https://web.ifzq.gtimg.cn/`
  - `https://proxy.finance.qq.com/`
  每个代理分别以 **HTTP 代理** 和 **HTTPS 代理** 各测一次；**协议** 写入 meta：只通 HTTP→`http`，只通 HTTPS→`https`，两个都通→`http/s`。去重按 **协议+IP+端口**。校验时**多代理并发**、**单代理内五域名并行**。列表直显：表格列为「代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议」。
- **更新频率**：列表按小时更新，保证可用代理的时效性。
- **并发参数**：校验 worker 数可通过 `-check-workers`（如 `-check-workers=4000`）或环境变量 `GFP_CHECK_WORKERS` 设置，默认 4000，最大 4000。遇目标站限流可适当调低。

### 工作流程

1. **拉取**：从 `sources/` 目录下配置的源（仅处理 `http.txt`、`https.txt`）拉取原始代理数据，支持动态 URL 及 Base64 等格式。
2. **解析与规范化**：将原始数据解析为标准代理格式（协议、IP、端口、认证等）。
3. **校验**：对 HTTP/HTTPS 代理通过上述验证与 2 秒超时规则进行筛选。
4. **去重与存储**：通过校验的代理去重后写入内存。
5. **生成列表**：按协议生成 `list/` 目录下的 `http.txt`、`https.txt`，并更新统计与 README 中的下载表格。

自动化由 GitHub Actions 执行：**全量流程**（抓取→解析→验证→生成列表）**每 6 小时**运行一次；**轻量复测**（对已有列表做连通性复测、剔除失效代理）**每 1 小时**运行一次。全量任务最长运行 12 小时，超时才会取消。下表「最后更新」时间为 UTC 及 UTC+8。

### 支持的代理格式示例

| 类型 | 格式 | 示例 |
| :--- | :--- | :--- |
| **HTTP/HTTPS** | `http://ip:port` | `http://1.2.3.4:8080` |
| | `https://ip:port` | `https://1.2.3.4:8080` |
| | `http://user:pass@ip:port` | `http://user:pass@1.2.3.4:8080` |

---

## 🔗 直接下载链接

点击下方表格中您需要的协议类型即可获取最新列表，链接始终指向最近更新的代理文件。上述三个文件（http.txt、https.txt、passed.txt）同时发布至 [GitHub Releases (tag: lists)](https://github.com/feitianyul/free-proxy-list/releases/tag/lists)，每次运行覆盖同版本附件，可固定使用该 Release 的附件 URL。

<!-- BEGIN PROXY LIST -->

最后更新：2026-04-07 23:43:03 UTC（2026-04-08 07:43:03 UTC+8）

**代理总数：278**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 278 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 278 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 35.225.22.61:80 | 否 | ✓ 1387ms | ✓ 980ms | ✓ 1271ms | ✓ 963ms | http |
| 38.145.218.113:8446 | ✓ 433ms | ✓ 681ms | ✓ 1319ms | ✓ 1915ms | ✓ 581ms | http |
| 147.161.210.140:8800 | ✓ 540ms | ✓ 1534ms | ✓ 584ms | ✓ 1181ms | ✓ 946ms | http |
| 152.32.132.190:7890 | ✓ 641ms | ✓ 1752ms | ✓ 632ms | ✓ 983ms | ✓ 634ms | http |
| 178.128.24.162:8080 | 否 | 否 | ✓ 1115ms | ✓ 1041ms | ✓ 791ms | http |
| 218.60.175.168:22222 | ✓ 887ms | ✓ 1230ms | ✓ 995ms | ✓ 1227ms | ✓ 977ms | http |
| 111.227.254.10:22222 | ✓ 964ms | ✓ 1258ms | ✓ 913ms | ✓ 1316ms | ✓ 1065ms | http |
| 111.227.254.11:22222 | ✓ 956ms | ✓ 1285ms | ✓ 955ms | ✓ 1237ms | ✓ 1074ms | http |
| 159.223.71.162:8080 | ✓ 863ms | 否 | ✓ 1115ms | ✓ 1351ms | ✓ 797ms | http |
| 159.223.71.162:443 | ✓ 862ms | 否 | ✓ 1114ms | ✓ 1357ms | ✓ 793ms | http |
| 111.227.254.12:22222 | ✓ 996ms | ✓ 1321ms | ✓ 996ms | ✓ 1305ms | ✓ 1045ms | http |
| 111.227.254.9:22222 | ✓ 1005ms | ✓ 1299ms | ✓ 1008ms | ✓ 1320ms | ✓ 1060ms | http |
| 1.231.81.166:3128 | ✓ 687ms | ✓ 1877ms | ✓ 1813ms | ✓ 1176ms | ✓ 1119ms | http |
| 167.103.115.102:8800 | ✓ 1075ms | 否 | ✓ 1111ms | ✓ 1104ms | ✓ 1183ms | http |
| 113.160.132.26:8080 | ✓ 1359ms | ✓ 1302ms | ✓ 1107ms | ✓ 1524ms | ✓ 1254ms | http |
| 147.161.239.240:8800 | ✓ 819ms | ✓ 1911ms | ✓ 1561ms | ✓ 1489ms | ✓ 1586ms | http |
| 161.35.70.36:8888 | ✓ 788ms | 否 | ✓ 1505ms | 否 | ✓ 1475ms | http |
| 1.225.116.115:1080 | 否 | ✓ 1523ms | ✓ 1616ms | ✓ 1686ms | ✓ 928ms | http |
| 5.104.87.17:8051 | 否 | 否 | ✓ 1400ms | ✓ 1647ms | ✓ 1475ms | http |
| 159.223.213.91:8888 | ✓ 865ms | ✓ 1971ms | ✓ 1896ms | 否 | ✓ 1547ms | http |
| 167.103.34.108:8800 | ✓ 1903ms | ✓ 1965ms | ✓ 1536ms | ✓ 1598ms | ✓ 1606ms | http |
| 165.22.57.158:8080 | ✓ 866ms | ✓ 1930ms | ✓ 1231ms | 否 | 否 | http |
| 109.107.179.140:8090 | ✓ 1004ms | 否 | ✓ 1346ms | 否 | ✓ 1681ms | http |
| 45.167.124.52:8080 | 否 | ✓ 1834ms | 否 | ✓ 1982ms | ✓ 1686ms | http |
| 38.34.179.19:8448 | ✓ 668ms | ✓ 766ms | ✓ 79ms | ✓ 639ms | ✓ 511ms | http |
| 38.34.179.201:8444 | ✓ 582ms | ✓ 659ms | ✓ 98ms | ✓ 644ms | ✓ 620ms | http |
| 38.34.179.7:8444 | ✓ 621ms | ✓ 801ms | ✓ 75ms | ✓ 645ms | ✓ 498ms | http |
| 38.34.179.229:8447 | ✓ 671ms | ✓ 649ms | ✓ 98ms | ✓ 705ms | ✓ 575ms | http |
| 38.34.179.204:8447 | ✓ 618ms | ✓ 625ms | ✓ 124ms | ✓ 703ms | ✓ 582ms | http |
| 38.145.203.108:8452 | ✓ 599ms | ✓ 609ms | ✓ 139ms | ✓ 706ms | ✓ 631ms | http |
| 38.145.203.98:8446 | ✓ 609ms | ✓ 640ms | ✓ 106ms | ✓ 703ms | ✓ 647ms | http |
| 38.145.203.97:8446 | ✓ 637ms | ✓ 639ms | ✓ 116ms | ✓ 697ms | ✓ 645ms | http |
| 38.34.179.66:8444 | ✓ 591ms | ✓ 783ms | ✓ 94ms | ✓ 666ms | ✓ 520ms | http |
| 45.136.130.193:8447 | ✓ 613ms | ✓ 595ms | ✓ 152ms | ✓ 685ms | ✓ 674ms | http |
| 45.136.130.188:8449 | ✓ 632ms | ✓ 614ms | ✓ 151ms | ✓ 671ms | ✓ 668ms | http |
| 45.136.131.47:8452 | ✓ 666ms | ✓ 801ms | ✓ 80ms | ✓ 676ms | ✓ 587ms | http |
| 45.136.131.54:8448 | ✓ 600ms | ✓ 845ms | ✓ 80ms | ✓ 682ms | ✓ 535ms | http |
| 45.136.131.57:8449 | ✓ 583ms | ✓ 749ms | ✓ 84ms | ✓ 666ms | ✓ 647ms | http |
| 38.145.220.49:8444 | ✓ 665ms | ✓ 850ms | ✓ 93ms | ✓ 715ms | ✓ 490ms | http |
| 38.34.183.224:8448 | ✓ 584ms | ✓ 799ms | ✓ 79ms | ✓ 794ms | ✓ 485ms | http |
| 38.145.218.238:8443 | ✓ 589ms | ✓ 835ms | ✓ 74ms | ✓ 763ms | ✓ 490ms | http |
| 38.145.220.22:8443 | ✓ 649ms | ✓ 797ms | ✓ 84ms | ✓ 764ms | ✓ 511ms | http |
| 38.34.179.67:8446 | ✓ 601ms | ✓ 821ms | ✓ 78ms | ✓ 707ms | ✓ 526ms | http |
| 38.145.220.103:8446 | ✓ 669ms | ✓ 746ms | ✓ 80ms | ✓ 714ms | ✓ 627ms | http |
| 38.34.183.47:8452 | ✓ 602ms | ✓ 768ms | ✓ 152ms | ✓ 758ms | ✓ 487ms | http |
| 38.145.203.86:8449 | ✓ 598ms | ✓ 742ms | ✓ 78ms | ✓ 740ms | ✓ 620ms | http |
| 38.145.220.43:8450 | ✓ 655ms | ✓ 862ms | ✓ 75ms | ✓ 735ms | ✓ 511ms | http |
| 38.34.183.13:8449 | ✓ 611ms | ✓ 770ms | ✓ 91ms | ✓ 814ms | ✓ 524ms | http |
| 38.34.183.11:8446 | ✓ 667ms | ✓ 787ms | ✓ 92ms | ✓ 806ms | ✓ 513ms | http |
| 38.145.218.101:8450 | ✓ 584ms | ✓ 770ms | ✓ 102ms | ✓ 781ms | ✓ 539ms | http |
| 38.145.208.227:8447 | ✓ 627ms | ✓ 799ms | ✓ 91ms | ✓ 811ms | ✓ 514ms | http |
| 38.145.218.206:8444 | ✓ 586ms | ✓ 779ms | ✓ 80ms | ✓ 869ms | ✓ 509ms | http |
| 38.145.220.39:8452 | ✓ 583ms | ✓ 786ms | ✓ 79ms | ✓ 866ms | ✓ 508ms | http |
| 38.145.220.55:8443 | ✓ 616ms | ✓ 778ms | ✓ 171ms | ✓ 782ms | ✓ 515ms | http |
| 38.145.220.40:8450 | ✓ 588ms | ✓ 818ms | ✓ 74ms | ✓ 838ms | ✓ 519ms | http |
| 38.145.220.77:8453 | ✓ 689ms | ✓ 674ms | ✓ 100ms | ✓ 856ms | ✓ 510ms | http |
| 38.34.183.221:8452 | ✓ 636ms | ✓ 838ms | ✓ 97ms | ✓ 737ms | ✓ 653ms | http |
| 38.34.179.155:8453 | ✓ 659ms | ✓ 777ms | ✓ 90ms | ✓ 813ms | ✓ 668ms | http |
| 38.34.179.156:8451 | ✓ 602ms | ✓ 768ms | ✓ 88ms | ✓ 842ms | ✓ 823ms | http |
| 45.136.131.25:8453 | ✓ 653ms | ✓ 592ms | ✓ 576ms | ✓ 724ms | ✓ 660ms | http |
| 45.136.131.26:8446 | ✓ 622ms | ✓ 634ms | ✓ 532ms | ✓ 761ms | ✓ 640ms | http |
| 45.136.131.35:8452 | ✓ 611ms | ✓ 639ms | ✓ 532ms | ✓ 769ms | ✓ 628ms | http |
| 45.136.131.28:8449 | ✓ 595ms | ✓ 647ms | ✓ 516ms | ✓ 751ms | ✓ 650ms | http |
| 45.136.131.33:8452 | ✓ 650ms | ✓ 641ms | ✓ 530ms | ✓ 766ms | ✓ 631ms | http |
| 45.136.131.58:8445 | ✓ 585ms | ✓ 828ms | ✓ 412ms | ✓ 748ms | ✓ 620ms | http |
| 45.136.131.67:8448 | ✓ 628ms | ✓ 824ms | ✓ 402ms | ✓ 740ms | ✓ 647ms | http |
| 45.136.131.61:8444 | ✓ 659ms | ✓ 822ms | ✓ 408ms | ✓ 752ms | ✓ 633ms | http |
| 38.145.208.209:8447 | ✓ 663ms | ✓ 607ms | ✓ 697ms | ✓ 834ms | ✓ 507ms | http |
| 38.34.179.94:8444 | ✓ 629ms | ✓ 777ms | ✓ 679ms | ✓ 691ms | ✓ 522ms | http |
| 38.145.208.211:8453 | ✓ 981ms | ✓ 595ms | ✓ 352ms | ✓ 852ms | ✓ 512ms | http |
| 38.145.208.208:8444 | ✓ 952ms | ✓ 673ms | ✓ 270ms | ✓ 839ms | ✓ 533ms | http |
| 38.34.179.87:8447 | ✓ 1027ms | ✓ 680ms | ✓ 347ms | ✓ 757ms | ✓ 535ms | http |
| 38.145.208.203:8449 | ✓ 627ms | ✓ 635ms | ✓ 658ms | ✓ 851ms | ✓ 530ms | http |
| 38.34.179.39:8452 | ✓ 650ms | ✓ 845ms | ✓ 595ms | ✓ 795ms | ✓ 494ms | http |
| 38.34.179.36:8453 | ✓ 591ms | ✓ 777ms | ✓ 669ms | ✓ 794ms | ✓ 516ms | http |
| 38.34.179.38:8447 | ✓ 586ms | ✓ 802ms | ✓ 650ms | ✓ 816ms | ✓ 494ms | http |
| 38.34.179.40:8446 | ✓ 613ms | ✓ 764ms | ✓ 680ms | ✓ 822ms | ✓ 512ms | http |
| 38.34.179.165:8450 | ✓ 1001ms | ✓ 639ms | ✓ 425ms | ✓ 687ms | ✓ 671ms | http |
| 38.34.179.176:8446 | ✓ 987ms | ✓ 599ms | ✓ 460ms | ✓ 702ms | ✓ 670ms | http |
| 38.34.179.177:8446 | ✓ 955ms | ✓ 596ms | ✓ 466ms | ✓ 683ms | ✓ 693ms | http |
| 38.145.208.173:8444 | ✓ 1014ms | ✓ 605ms | ✓ 574ms | ✓ 825ms | ✓ 509ms | http |
| 38.145.208.174:8444 | ✓ 975ms | ✓ 609ms | ✓ 573ms | ✓ 825ms | ✓ 506ms | http |
| 38.145.208.169:8452 | ✓ 992ms | ✓ 679ms | ✓ 496ms | ✓ 848ms | ✓ 519ms | http |
| 38.145.208.246:8450 | ✓ 959ms | ✓ 797ms | ✓ 468ms | ✓ 793ms | ✓ 533ms | http |
| 38.145.208.241:8447 | ✓ 1014ms | ✓ 793ms | ✓ 468ms | ✓ 818ms | ✓ 521ms | http |
| 38.145.208.243:8447 | ✓ 1034ms | ✓ 797ms | ✓ 463ms | ✓ 850ms | ✓ 499ms | http |
| 38.145.208.244:8446 | ✓ 954ms | ✓ 814ms | ✓ 449ms | ✓ 849ms | ✓ 502ms | http |
| 38.145.208.242:8444 | ✓ 966ms | ✓ 841ms | ✓ 449ms | ✓ 823ms | ✓ 507ms | http |
| 38.145.208.247:8452 | ✓ 979ms | ✓ 797ms | ✓ 462ms | ✓ 846ms | ✓ 517ms | http |
| 38.34.179.51:8449 | ✓ 996ms | ✓ 769ms | ✓ 454ms | ✓ 690ms | ✓ 551ms | http |
| 38.34.179.101:8453 | ✓ 1015ms | ✓ 600ms | ✓ 468ms | ✓ 811ms | ✓ 731ms | http |
| 38.34.179.57:8448 | ✓ 1038ms | ✓ 812ms | ✓ 413ms | ✓ 687ms | ✓ 539ms | http |
| 38.34.179.61:8445 | ✓ 982ms | ✓ 597ms | ✓ 517ms | ✓ 786ms | ✓ 711ms | http |
| 38.145.218.216:8449 | ✓ 593ms | ✓ 638ms | ✓ 110ms | ✓ 709ms | ✓ 651ms | http |
| 38.34.179.24:8447 | ✓ 1001ms | ✓ 619ms | ✓ 805ms | ✓ 826ms | ✓ 491ms | http |
| 38.34.179.35:8443 | ✓ 620ms | 否 | ✓ 193ms | ✓ 671ms | ✓ 515ms | http |
| 38.145.208.204:8446 | ✓ 615ms | ✓ 744ms | ✓ 100ms | ✓ 779ms | ✓ 823ms | http |
| 38.34.179.77:8448 | ✓ 622ms | ✓ 820ms | ✓ 93ms | ✓ 720ms | ✓ 794ms | http |
| 38.34.179.78:8448 | ✓ 594ms | ✓ 861ms | ✓ 98ms | ✓ 766ms | ✓ 708ms | http |
| 38.145.203.34:8444 | ✓ 985ms | ✓ 606ms | ✓ 170ms | ✓ 1355ms | ✓ 1933ms | http |

<!-- END PROXY TABLE -->

## 🤝 参与贡献

本项目由社区驱动，欢迎任何形式的贡献。最简单的参与方式就是添加新的代理数据源。

请先阅读 **[贡献指南](CONTRIBUTING.md)** 了解如何开始。

## 🙏 支持本项目

如果您觉得本项目有帮助，欢迎给予支持，让更多人看到并参与贡献。

-   在 GitHub 上 **给本仓库加星** ⭐️
-   **分享**给朋友和同事

## ⚠️ 免责声明

-   本仓库中的代理均来自公开来源，不保证其速度、安全性或可用性。
-   使用这些代理的风险由您自行承担。
-   本仓库维护者不对任何滥用行为负责。请勿将代理用于非法用途。

## 📝 许可证

本仓库采用 MIT 许可证发布。详见 [LICENSE](LICENSE)。

## Stars
[![Star History Chart](https://api.star-history.com/svg?repos=feitianyul/free-proxy-list&type=Date)](https://star-history.com/#feitianyul/free-proxy-list&Date)
