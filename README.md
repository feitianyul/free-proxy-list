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

最后更新：2026-04-02 23:29:34 UTC（2026-04-03 07:29:34 UTC+8）

**代理总数：143**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 143 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 143 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 38.34.179.61:8445 | ✓ 404ms | ✓ 892ms | ✓ 364ms | ✓ 936ms | ✓ 780ms | http |
| 38.145.220.102:8453 | ✓ 1250ms | ✓ 1260ms | ✓ 95ms | ✓ 699ms | ✓ 733ms | http |
| 147.161.210.140:8800 | ✓ 1182ms | ✓ 1123ms | ✓ 788ms | ✓ 920ms | ✓ 1102ms | http |
| 38.145.208.242:8451 | ✓ 889ms | ✓ 787ms | ✓ 460ms | ✓ 1544ms | ✓ 1633ms | http |
| 5.104.87.17:8051 | ✓ 1137ms | 否 | ✓ 890ms | ✓ 915ms | ✓ 997ms | http |
| 115.231.181.40:8128 | ✓ 829ms | ✓ 1401ms | ✓ 967ms | ✓ 1248ms | ✓ 890ms | http |
| 159.223.71.162:8080 | ✓ 734ms | 否 | ✓ 1104ms | ✓ 1062ms | ✓ 1105ms | http |
| 1.231.81.166:3128 | ✓ 1214ms | ✓ 1499ms | ✓ 1482ms | ✓ 1015ms | ✓ 1344ms | http |
| 159.223.71.162:443 | ✓ 745ms | 否 | ✓ 1110ms | ✓ 1125ms | ✓ 1450ms | http |
| 167.103.115.102:8800 | ✓ 1132ms | ✓ 1671ms | ✓ 967ms | ✓ 1066ms | ✓ 1966ms | http |
| 113.160.132.26:8080 | ✓ 1544ms | ✓ 1391ms | ✓ 886ms | ✓ 1386ms | ✓ 1353ms | http |
| 203.80.138.81:50000 | ✓ 884ms | ✓ 1122ms | 否 | ✓ 1019ms | ✓ 799ms | http |
| 95.213.217.168:52004 | ✓ 1065ms | ✓ 1751ms | 否 | 否 | ✓ 1718ms | http |
| 167.103.34.108:8800 | ✓ 1564ms | ✓ 1996ms | ✓ 1455ms | ✓ 1818ms | ✓ 1520ms | http |
| 42.96.16.158:1311 | ✓ 1533ms | 否 | ✓ 1241ms | ✓ 1203ms | ✓ 936ms | http |
| 45.167.124.52:8080 | ✓ 1128ms | 否 | ✓ 1543ms | 否 | ✓ 1504ms | http |
| 180.250.219.58:53281 | ✓ 1547ms | 否 | ✓ 1520ms | 否 | ✓ 1901ms | http |
| 35.225.22.61:80 | ✓ 593ms | 否 | ✓ 514ms | ✓ 1144ms | ✓ 910ms | http |
| 38.34.179.27:8451 | ✓ 224ms | ✓ 633ms | ✓ 442ms | ✓ 1659ms | ✓ 1005ms | http |
| 38.34.179.98:8451 | ✓ 150ms | ✓ 654ms | ✓ 145ms | ✓ 732ms | ✓ 1447ms | http |
| 38.34.179.105:8449 | ✓ 598ms | ✓ 1691ms | ✓ 206ms | ✓ 806ms | ✓ 648ms | http |
| 38.34.183.47:8452 | ✓ 1194ms | ✓ 640ms | ✓ 262ms | ✓ 922ms | ✓ 1082ms | http |
| 38.34.183.13:8449 | ✓ 390ms | ✓ 658ms | ✓ 711ms | ✓ 824ms | ✓ 1186ms | http |
| 38.34.179.86:8452 | ✓ 644ms | ✓ 775ms | ✓ 727ms | ✓ 821ms | ✓ 1186ms | http |
| 43.99.54.236:5555 | ✓ 689ms | ✓ 902ms | ✓ 687ms | ✓ 830ms | ✓ 642ms | http |
| 38.34.179.14:8450 | ✓ 1342ms | ✓ 675ms | ✓ 118ms | ✓ 667ms | ✓ 931ms | http |
| 38.34.183.130:8452 | ✓ 1650ms | ✓ 653ms | ✓ 648ms | ✓ 939ms | ✓ 1830ms | http |
| 38.34.179.40:8446 | 否 | ✓ 1049ms | ✓ 252ms | ✓ 926ms | ✓ 1274ms | http |
| 38.34.183.222:8453 | 否 | ✓ 699ms | ✓ 406ms | ✓ 850ms | ✓ 1412ms | http |
| 38.34.179.173:8452 | 否 | ✓ 912ms | ✓ 104ms | ✓ 676ms | ✓ 637ms | http |
| 38.34.183.164:8444 | ✓ 1680ms | ✓ 649ms | ✓ 571ms | ✓ 875ms | ✓ 1623ms | http |
| 38.34.179.39:8452 | ✓ 1344ms | ✓ 1769ms | ✓ 250ms | ✓ 1195ms | 否 | http |
| 45.136.130.169:8444 | ✓ 1473ms | ✓ 779ms | ✓ 671ms | ✓ 1776ms | ✓ 1544ms | http |
| 34.96.238.40:8080 | ✓ 1198ms | 否 | ✓ 1123ms | ✓ 982ms | ✓ 1074ms | http |
| 38.34.179.150:8449 | ✓ 262ms | ✓ 789ms | ✓ 1713ms | ✓ 1307ms | ✓ 687ms | http |
| 38.34.183.225:8450 | ✓ 1649ms | ✓ 1254ms | ✓ 724ms | ✓ 1990ms | ✓ 1108ms | http |
| 38.34.179.54:8453 | ✓ 231ms | ✓ 618ms | ✓ 1150ms | 否 | ✓ 546ms | http |
| 167.103.144.127:8800 | ✓ 1562ms | 否 | ✓ 1343ms | ✓ 1563ms | 否 | http |
| 14.143.222.113:10155 | ✓ 959ms | 否 | ✓ 1202ms | ✓ 1679ms | 否 | http |
| 212.58.132.5:8888 | ✓ 1192ms | 否 | ✓ 1555ms | ✓ 1586ms | ✓ 1242ms | http |
| 38.145.208.207:8445 | ✓ 1087ms | ✓ 1692ms | ✓ 629ms | ✓ 1239ms | 否 | http |
| 120.92.212.16:8890 | ✓ 923ms | ✓ 1185ms | 否 | ✓ 1235ms | 否 | http |
| 38.34.183.224:8448 | ✓ 1828ms | ✓ 1198ms | ✓ 511ms | 否 | ✓ 1410ms | http |
| 45.167.125.21:999 | ✓ 1720ms | ✓ 1726ms | ✓ 1408ms | ✓ 1786ms | ✓ 1539ms | http |
| 208.87.243.199:7878 | ✓ 286ms | ✓ 815ms | ✓ 704ms | ✓ 881ms | ✓ 818ms | http |
| 165.232.146.249:3128 | ✓ 326ms | ✓ 797ms | ✓ 1298ms | ✓ 832ms | ✓ 739ms | http |
| 20.120.225.109:3128 | ✓ 1061ms | ✓ 1207ms | ✓ 816ms | ✓ 1858ms | ✓ 1059ms | http |
| 133.242.138.34:8100 | ✓ 1398ms | ✓ 1286ms | ✓ 1442ms | ✓ 1966ms | ✓ 961ms | http |
| 45.12.151.226:2829 | ✓ 1189ms | ✓ 1916ms | 否 | ✓ 1770ms | ✓ 1372ms | http |
| 59.46.216.131:30001 | ✓ 924ms | ✓ 1275ms | ✓ 1138ms | ✓ 1350ms | 否 | http |
| 120.92.212.16:7890 | ✓ 956ms | ✓ 1196ms | ✓ 999ms | 否 | 否 | http |
| 38.34.179.66:8444 | ✓ 342ms | ✓ 823ms | ✓ 392ms | ✓ 1007ms | ✓ 606ms | http |
| 38.34.179.186:8444 | ✓ 263ms | ✓ 612ms | ✓ 613ms | ✓ 1965ms | 否 | http |
| 45.136.131.29:8450 | ✓ 770ms | ✓ 625ms | ✓ 1403ms | ✓ 1032ms | ✓ 583ms | http |
| 128.199.121.61:9090 | ✓ 739ms | 否 | ✓ 795ms | ✓ 1047ms | 否 | http |
| 146.190.80.158:9090 | ✓ 730ms | 否 | 否 | ✓ 1263ms | ✓ 861ms | http |
| 101.43.127.100:8877 | ✓ 917ms | ✓ 1121ms | ✓ 895ms | ✓ 1097ms | ✓ 865ms | http |
| 128.199.116.219:9090 | ✓ 729ms | 否 | ✓ 956ms | ✓ 1100ms | ✓ 869ms | http |
| 147.161.239.240:8800 | ✓ 1163ms | ✓ 1719ms | ✓ 1340ms | ✓ 1715ms | 否 | http |
| 47.105.98.23:3128 | 否 | ✓ 1154ms | 否 | ✓ 1218ms | ✓ 1481ms | http |
| 38.145.218.102:8447 | ✓ 794ms | ✓ 760ms | ✓ 1400ms | 否 | ✓ 1909ms | http |
| 128.199.114.189:9090 | ✓ 906ms | 否 | ✓ 1162ms | ✓ 1409ms | ✓ 830ms | http |
| 177.234.217.88:999 | ✓ 1208ms | ✓ 1841ms | ✓ 1752ms | 否 | ✓ 1842ms | http |
| 38.145.208.210:8448 | ✓ 1525ms | ✓ 1854ms | ✓ 1973ms | 否 | 否 | http |
| 45.140.147.155:1082 | ✓ 1758ms | ✓ 1531ms | ✓ 1296ms | ✓ 1865ms | ✓ 1259ms | http |
| 45.140.147.155:1081 | ✓ 1070ms | 否 | ✓ 1523ms | ✓ 1849ms | ✓ 1264ms | http |
| 116.80.65.85:3172 | ✓ 1490ms | ✓ 1955ms | 否 | 否 | ✓ 1627ms | http |
| 128.199.113.85:9090 | ✓ 1353ms | 否 | ✓ 788ms | ✓ 1056ms | ✓ 897ms | http |
| 62.171.161.88:2018 | ✓ 624ms | ✓ 1949ms | 否 | ✓ 1930ms | ✓ 1516ms | http |
| 167.103.31.122:8800 | ✓ 1421ms | 否 | ✓ 1325ms | 否 | ✓ 1588ms | http |
| 190.12.150.244:999 | ✓ 1402ms | 否 | ✓ 891ms | 否 | ✓ 1440ms | http |
| 195.123.209.48:3128 | ✓ 1171ms | 否 | ✓ 1564ms | 否 | ✓ 1672ms | http |
| 82.114.228.67:1080 | ✓ 1399ms | ✓ 1694ms | 否 | ✓ 1729ms | 否 | http |
| 83.219.250.8:62920 | ✓ 1045ms | 否 | ✓ 1061ms | ✓ 1732ms | ✓ 1500ms | http |
| 168.110.52.228:3128 | ✓ 462ms | 否 | ✓ 452ms | ✓ 766ms | 否 | http |
| 175.194.173.105:3128 | 否 | 否 | ✓ 809ms | ✓ 1035ms | ✓ 905ms | http |
| 72.11.151.159:6005 | ✓ 633ms | ✓ 1387ms | ✓ 1004ms | ✓ 1495ms | ✓ 1207ms | http |
| 34.101.184.164:3128 | ✓ 1056ms | 否 | ✓ 1039ms | ✓ 1441ms | ✓ 1153ms | http |
| 45.136.131.36:8450 | ✓ 642ms | ✓ 630ms | ✓ 528ms | 否 | ✓ 1294ms | http |
| 111.79.111.126:3128 | ✓ 1087ms | 否 | ✓ 1422ms | 否 | ✓ 1919ms | http |
| 121.232.73.214:1080 | ✓ 966ms | ✓ 1098ms | ✓ 1031ms | ✓ 1332ms | ✓ 1084ms | http |
| 38.145.208.172:8448 | ✓ 678ms | ✓ 874ms | ✓ 152ms | ✓ 681ms | ✓ 537ms | http |
| 38.145.218.232:8448 | ✓ 701ms | ✓ 636ms | ✓ 366ms | ✓ 795ms | ✓ 636ms | http |
| 38.34.183.234:8450 | ✓ 690ms | ✓ 918ms | ✓ 251ms | ✓ 671ms | ✓ 558ms | http |
| 38.145.218.212:8448 | ✓ 700ms | ✓ 633ms | ✓ 370ms | ✓ 792ms | ✓ 650ms | http |
| 38.145.220.33:8448 | ✓ 698ms | ✓ 667ms | ✓ 338ms | ✓ 736ms | ✓ 785ms | http |
| 38.145.220.39:8447 | ✓ 699ms | ✓ 881ms | ✓ 488ms | ✓ 988ms | ✓ 713ms | http |
| 160.250.5.22:1 | ✓ 1661ms | 否 | 否 | ✓ 1247ms | ✓ 1024ms | http |
| 128.199.254.13:9090 | ✓ 842ms | 否 | ✓ 1115ms | ✓ 1073ms | ✓ 891ms | http |
| 38.34.179.21:8446 | ✓ 128ms | ✓ 615ms | ✓ 95ms | ✓ 710ms | ✓ 513ms | http |
| 38.34.179.174:8453 | ✓ 180ms | ✓ 669ms | ✓ 171ms | ✓ 729ms | ✓ 760ms | http |
| 38.34.179.193:8452 | ✓ 300ms | ✓ 702ms | ✓ 121ms | ✓ 824ms | ✓ 509ms | http |
| 38.34.179.51:8449 | ✓ 248ms | ✓ 758ms | ✓ 132ms | ✓ 694ms | ✓ 531ms | http |
| 38.34.179.8:8444 | ✓ 432ms | ✓ 856ms | ✓ 386ms | ✓ 650ms | ✓ 522ms | http |
| 38.145.220.11:8447 | ✓ 851ms | ✓ 1701ms | ✓ 106ms | ✓ 728ms | ✓ 756ms | http |
| 209.126.84.232:8888 | ✓ 855ms | ✓ 1981ms | ✓ 955ms | 否 | ✓ 1889ms | http |
| 45.136.130.186:8451 | ✓ 404ms | ✓ 1530ms | ✓ 1673ms | ✓ 1371ms | 否 | http |
| 38.145.218.101:8447 | ✓ 1295ms | 否 | ✓ 1628ms | ✓ 1786ms | 否 | http |
| 103.22.217.30:8097 | ✓ 1239ms | 否 | ✓ 1354ms | ✓ 1630ms | ✓ 1388ms | http |
| 38.180.2.107:3128 | ✓ 944ms | ✓ 1768ms | ✓ 1824ms | 否 | ✓ 1698ms | http |

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
