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

最后更新：2026-03-29 19:45:38 UTC（2026-03-30 03:45:38 UTC+8）

**代理总数：189**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 188 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 189 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 39.185.46.193:5911 | ✓ 580ms | ✓ 732ms | ✓ 585ms | ✓ 958ms | ✓ 626ms | http |
| 35.225.22.61:80 | ✓ 785ms | 否 | ✓ 612ms | ✓ 1281ms | 否 | http |
| 147.161.210.140:8800 | ✓ 1374ms | 否 | ✓ 864ms | ✓ 693ms | ✓ 756ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 680ms | ✓ 854ms | ✓ 675ms | http |
| 167.103.115.102:8800 | ✓ 816ms | 否 | ✓ 923ms | ✓ 1507ms | ✓ 906ms | http |
| 42.96.16.158:1311 | ✓ 1301ms | 否 | ✓ 1131ms | ✓ 1128ms | ✓ 893ms | http |
| 147.161.239.240:8800 | ✓ 1238ms | ✓ 1902ms | ✓ 1323ms | ✓ 1463ms | ✓ 1270ms | http |
| 95.213.217.168:52004 | ✓ 1270ms | 否 | ✓ 1422ms | ✓ 1807ms | ✓ 1339ms | http |
| 167.103.34.108:8800 | ✓ 1426ms | 否 | ✓ 1538ms | ✓ 1544ms | ✓ 1406ms | http |
| 45.167.124.52:8080 | 否 | ✓ 1934ms | 否 | ✓ 1889ms | ✓ 1614ms | http |
| 43.99.54.236:5555 | ✓ 1021ms | ✓ 899ms | ✓ 619ms | ✓ 758ms | 否 | http |
| 183.249.5.117:22222 | ✓ 646ms | ✓ 824ms | ✓ 613ms | ✓ 976ms | ✓ 662ms | http |
| 38.145.220.102:8453 | ✓ 994ms | ✓ 840ms | ✓ 490ms | ✓ 1118ms | ✓ 1717ms | http |
| 38.145.218.228:8447 | ✓ 490ms | ✓ 915ms | ✓ 950ms | ✓ 1088ms | ✓ 976ms | http |
| 222.184.48.242:22222 | ✓ 884ms | ✓ 1067ms | ✓ 869ms | ✓ 1023ms | ✓ 857ms | http |
| 113.160.132.26:8080 | ✓ 829ms | ✓ 1364ms | 否 | ✓ 1635ms | ✓ 929ms | http |
| 167.103.144.127:8800 | ✓ 1556ms | 否 | 否 | ✓ 1655ms | ✓ 1336ms | http |
| 167.103.31.122:8800 | ✓ 1816ms | 否 | ✓ 1703ms | ✓ 1768ms | 否 | http |
| 180.250.219.58:53281 | ✓ 1703ms | 否 | ✓ 1397ms | ✓ 1839ms | ✓ 1860ms | http |
| 208.87.243.199:7878 | ✓ 999ms | ✓ 738ms | ✓ 803ms | ✓ 781ms | ✓ 729ms | http |
| 35.206.88.200:8888 | 否 | ✓ 1294ms | ✓ 926ms | ✓ 1485ms | ✓ 1054ms | http |
| 209.64.46.210:3128 | ✓ 915ms | 否 | ✓ 675ms | ✓ 887ms | ✓ 995ms | http |
| 38.145.218.10:8450 | ✓ 902ms | ✓ 913ms | ✓ 596ms | ✓ 1939ms | ✓ 973ms | http |
| 103.133.254.4:3128 | ✓ 1690ms | 否 | ✓ 1648ms | ✓ 1757ms | ✓ 1394ms | http |
| 45.12.151.226:2829 | ✓ 1303ms | ✓ 1866ms | ✓ 861ms | ✓ 1885ms | 否 | http |
| 38.34.179.164:8448 | ✓ 1686ms | ✓ 1077ms | 否 | ✓ 1568ms | ✓ 734ms | http |
| 190.12.150.244:999 | ✓ 1410ms | ✓ 1973ms | ✓ 1835ms | 否 | ✓ 1757ms | http |
| 116.80.48.16:7777 | 否 | ✓ 1960ms | 否 | ✓ 1759ms | ✓ 1568ms | http |
| 222.184.48.251:22222 | ✓ 867ms | ✓ 1070ms | ✓ 884ms | ✓ 1132ms | ✓ 782ms | http |
| 5.104.87.17:8051 | ✓ 1344ms | 否 | ✓ 1332ms | ✓ 1192ms | ✓ 1012ms | http |
| 45.136.130.191:8453 | ✓ 542ms | ✓ 1026ms | ✓ 1131ms | 否 | 否 | http |
| 38.145.220.33:8448 | ✓ 661ms | ✓ 1076ms | ✓ 1158ms | 否 | 否 | http |
| 38.34.183.130:8452 | ✓ 1900ms | ✓ 993ms | ✓ 376ms | 否 | 否 | http |
| 38.34.179.162:8451 | ✓ 258ms | ✓ 1161ms | ✓ 901ms | ✓ 963ms | 否 | http |
| 38.145.208.206:8448 | ✓ 271ms | ✓ 977ms | ✓ 937ms | ✓ 1387ms | 否 | http |
| 38.145.208.205:8445 | ✓ 490ms | ✓ 1319ms | ✓ 1140ms | ✓ 774ms | 否 | http |
| 45.136.130.186:8451 | ✓ 342ms | ✓ 996ms | ✓ 1318ms | ✓ 1143ms | 否 | http |
| 38.34.179.16:8451 | ✓ 990ms | ✓ 1648ms | ✓ 254ms | ✓ 991ms | 否 | http |
| 38.34.183.11:8451 | ✓ 975ms | ✓ 862ms | ✓ 347ms | ✓ 781ms | ✓ 887ms | http |
| 38.34.183.8:8448 | ✓ 1143ms | ✓ 809ms | ✓ 338ms | ✓ 1002ms | 否 | http |
| 45.136.131.32:8445 | ✓ 615ms | ✓ 961ms | ✓ 246ms | ✓ 872ms | ✓ 754ms | http |
| 38.145.220.11:8446 | ✓ 511ms | ✓ 1231ms | ✓ 767ms | ✓ 805ms | ✓ 601ms | http |
| 38.34.183.234:8450 | ✓ 1019ms | ✓ 1433ms | ✓ 219ms | ✓ 817ms | 否 | http |
| 38.34.183.233:8448 | ✓ 971ms | ✓ 1629ms | ✓ 191ms | ✓ 788ms | 否 | http |
| 38.34.179.88:8446 | ✓ 544ms | ✓ 1056ms | ✓ 879ms | ✓ 864ms | 否 | http |
| 38.145.208.211:8453 | ✓ 253ms | ✓ 969ms | ✓ 713ms | ✓ 1328ms | ✓ 603ms | http |
| 38.145.208.209:8444 | ✓ 477ms | ✓ 1189ms | ✓ 1295ms | ✓ 768ms | 否 | http |
| 45.136.130.177:8448 | ✓ 418ms | ✓ 1110ms | ✓ 1127ms | ✓ 1172ms | 否 | http |
| 38.34.179.23:8444 | ✓ 674ms | ✓ 1140ms | ✓ 663ms | ✓ 1501ms | 否 | http |
| 38.34.179.25:8444 | ✓ 668ms | ✓ 1145ms | ✓ 683ms | ✓ 1522ms | 否 | http |
| 38.145.208.241:8453 | 否 | ✓ 1608ms | 否 | ✓ 819ms | ✓ 930ms | http |
| 45.136.131.38:8445 | ✓ 646ms | ✓ 796ms | ✓ 357ms | ✓ 759ms | ✓ 598ms | http |
| 38.145.208.216:8448 | ✓ 649ms | ✓ 714ms | ✓ 437ms | ✓ 937ms | ✓ 826ms | http |
| 101.43.127.100:8877 | ✓ 751ms | ✓ 942ms | ✓ 834ms | ✓ 963ms | ✓ 831ms | http |
| 120.92.212.16:7890 | ✓ 1538ms | ✓ 1125ms | ✓ 846ms | ✓ 1162ms | ✓ 897ms | http |
| 38.145.208.188:8447 | ✓ 674ms | ✓ 746ms | ✓ 1210ms | 否 | ✓ 647ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1163ms | ✓ 990ms | ✓ 1127ms | ✓ 917ms | http |
| 193.233.22.29:10808 | ✓ 542ms | 否 | 否 | ✓ 1408ms | ✓ 1208ms | http |
| 5.102.109.41:999 | 否 | ✓ 1448ms | ✓ 1759ms | ✓ 1560ms | ✓ 1296ms | http |
| 177.234.217.88:999 | ✓ 1460ms | ✓ 1842ms | ✓ 1940ms | 否 | ✓ 1771ms | http |
| 38.34.179.173:8450 | ✓ 935ms | ✓ 987ms | ✓ 380ms | ✓ 767ms | ✓ 705ms | http |
| 38.145.203.76:8446 | ✓ 924ms | ✓ 1177ms | ✓ 202ms | ✓ 1005ms | ✓ 1217ms | http |
| 82.146.58.184:1080 | ✓ 866ms | ✓ 1895ms | ✓ 1436ms | 否 | 否 | http |
| 183.249.5.110:22222 | ✓ 701ms | ✓ 781ms | ✓ 821ms | ✓ 860ms | ✓ 674ms | http |
| 183.249.5.111:22222 | ✓ 642ms | ✓ 816ms | ✓ 658ms | ✓ 856ms | ✓ 836ms | http |
| 47.74.226.8:5001 | ✓ 849ms | 否 | 否 | ✓ 1323ms | ✓ 1112ms | http |
| 222.184.48.252:22222 | ✓ 1073ms | ✓ 1100ms | ✓ 945ms | ✓ 1203ms | ✓ 1159ms | http |
| 101.47.73.135:3128 | ✓ 1142ms | 否 | ✓ 1573ms | ✓ 1117ms | 否 | http |
| 149.62.191.202:3128 | ✓ 1295ms | 否 | ✓ 1551ms | 否 | ✓ 1523ms | http |
| 222.228.171.92:8080 | 否 | ✓ 1920ms | ✓ 1721ms | ✓ 1286ms | ✓ 982ms | http |
| 114.237.77.244:1080 | 否 | ✓ 1158ms | ✓ 1614ms | 否 | ✓ 1862ms | http |
| 210.223.44.230:3128 | ✓ 1506ms | ✓ 905ms | ✓ 1881ms | ✓ 1904ms | 否 | http |
| 38.145.218.160:8450 | ✓ 366ms | ✓ 757ms | ✓ 509ms | 否 | 否 | http |
| 106.75.15.167:7890 | ✓ 1573ms | ✓ 1092ms | ✓ 936ms | 否 | ✓ 1118ms | http |
| 38.145.208.240:8452 | ✓ 391ms | ✓ 757ms | ✓ 183ms | ✓ 1029ms | ✓ 702ms | http |
| 38.34.179.76:8453 | ✓ 261ms | ✓ 867ms | ✓ 268ms | ✓ 1144ms | ✓ 596ms | http |
| 38.34.179.106:8446 | ✓ 273ms | ✓ 748ms | ✓ 247ms | ✓ 1094ms | ✓ 825ms | http |
| 38.34.179.102:8448 | ✓ 183ms | ✓ 832ms | ✓ 172ms | ✓ 779ms | ✓ 562ms | http |
| 38.34.179.14:8450 | ✓ 175ms | ✓ 986ms | ✓ 191ms | ✓ 789ms | ✓ 597ms | http |
| 38.34.179.8:8449 | ✓ 169ms | ✓ 982ms | ✓ 195ms | ✓ 762ms | ✓ 581ms | http |
| 38.34.179.19:8449 | ✓ 179ms | ✓ 1007ms | ✓ 187ms | ✓ 778ms | ✓ 593ms | http |
| 38.34.179.26:8450 | ✓ 234ms | ✓ 918ms | ✓ 168ms | ✓ 757ms | ✓ 708ms | http |
| 38.145.203.19:8449 | ✓ 518ms | ✓ 852ms | ✓ 185ms | ✓ 790ms | ✓ 614ms | http |
| 38.34.179.178:8445 | ✓ 240ms | ✓ 896ms | ✓ 190ms | ✓ 1002ms | ✓ 611ms | http |
| 38.145.203.135:8444 | ✓ 286ms | ✓ 1072ms | ✓ 191ms | ✓ 783ms | ✓ 588ms | http |
| 38.34.179.173:8452 | ✓ 254ms | ✓ 888ms | ✓ 335ms | ✓ 902ms | ✓ 602ms | http |
| 38.34.179.27:8451 | ✓ 402ms | ✓ 851ms | ✓ 171ms | ✓ 757ms | ✓ 784ms | http |
| 38.34.179.98:8451 | ✓ 181ms | ✓ 905ms | ✓ 173ms | ✓ 791ms | ✓ 606ms | http |
| 38.34.179.57:8453 | ✓ 201ms | ✓ 887ms | ✓ 181ms | ✓ 778ms | ✓ 608ms | http |
| 38.145.218.87:8445 | ✓ 839ms | ✓ 710ms | ✓ 176ms | ✓ 902ms | ✓ 607ms | http |
| 38.34.179.174:8453 | ✓ 352ms | ✓ 971ms | ✓ 370ms | ✓ 814ms | ✓ 580ms | http |
| 38.145.208.244:8448 | ✓ 252ms | ✓ 1131ms | ✓ 203ms | ✓ 773ms | ✓ 731ms | http |
| 38.34.179.186:8444 | ✓ 842ms | ✓ 721ms | ✓ 174ms | ✓ 789ms | ✓ 592ms | http |
| 38.34.179.39:8452 | ✓ 319ms | ✓ 1367ms | ✓ 197ms | ✓ 811ms | ✓ 640ms | http |
| 219.117.204.211:7799 | ✓ 1276ms | ✓ 1421ms | ✓ 487ms | ✓ 747ms | ✓ 577ms | http |
| 38.34.183.13:8449 | ✓ 178ms | ✓ 1006ms | ✓ 349ms | ✓ 992ms | ✓ 762ms | http |
| 38.34.179.6:8449 | ✓ 638ms | ✓ 704ms | ✓ 187ms | ✓ 819ms | ✓ 749ms | http |
| 45.136.131.36:8450 | ✓ 303ms | ✓ 886ms | ✓ 215ms | ✓ 890ms | ✓ 592ms | http |
| 45.136.131.62:8449 | ✓ 218ms | ✓ 933ms | ✓ 163ms | ✓ 780ms | ✓ 593ms | http |
| 45.136.131.54:8448 | ✓ 231ms | ✓ 747ms | ✓ 793ms | ✓ 804ms | ✓ 639ms | http |

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
