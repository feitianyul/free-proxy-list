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

最后更新：2026-03-23 23:27:27 UTC（2026-03-24 07:27:27 UTC+8）

**代理总数：136**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 136 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 136 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 8.217.106.71:8888 | ✓ 609ms | ✓ 871ms | ✓ 644ms | ✓ 774ms | ✓ 620ms | http |
| 38.145.208.213:8450 | ✓ 916ms | 否 | ✓ 696ms | ✓ 956ms | ✓ 1299ms | http |
| 147.161.210.140:8800 | ✓ 1412ms | ✓ 1060ms | ✓ 802ms | ✓ 923ms | ✓ 1001ms | http |
| 35.225.22.61:80 | ✓ 947ms | ✓ 1443ms | ✓ 1011ms | ✓ 1161ms | ✓ 1230ms | http |
| 13.230.197.159:10808 | ✓ 1411ms | ✓ 1160ms | ✓ 924ms | ✓ 936ms | ✓ 939ms | http |
| 219.117.204.211:7799 | ✓ 1412ms | ✓ 1087ms | ✓ 991ms | 否 | ✓ 883ms | http |
| 1.231.81.166:3128 | ✓ 1442ms | ✓ 909ms | 否 | ✓ 1354ms | ✓ 1061ms | http |
| 113.160.132.26:8080 | ✓ 1421ms | ✓ 1339ms | ✓ 1253ms | ✓ 1251ms | ✓ 1299ms | http |
| 167.103.34.108:8800 | ✓ 1682ms | 否 | ✓ 1800ms | ✓ 1692ms | ✓ 1463ms | http |
| 43.99.54.236:5555 | ✓ 621ms | ✓ 886ms | ✓ 624ms | ✓ 791ms | ✓ 631ms | http |
| 166.88.55.83:7890 | ✓ 600ms | ✓ 1034ms | ✓ 603ms | ✓ 744ms | ✓ 601ms | http |
| 150.249.255.91:3128 | 否 | ✓ 1121ms | ✓ 504ms | ✓ 893ms | 否 | http |
| 155.212.132.241:3128 | ✓ 792ms | 否 | ✓ 1014ms | ✓ 1901ms | ✓ 1483ms | http |
| 91.233.223.147:3128 | ✓ 957ms | 否 | ✓ 980ms | 否 | ✓ 1692ms | http |
| 120.92.212.16:7890 | ✓ 1634ms | 否 | 否 | ✓ 1159ms | ✓ 950ms | http |
| 167.103.31.122:8800 | ✓ 1948ms | 否 | ✓ 1661ms | 否 | ✓ 1847ms | http |
| 144.31.79.117:8888 | ✓ 864ms | 否 | ✓ 1765ms | 否 | ✓ 1941ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1516ms | ✓ 955ms | ✓ 1198ms | ✓ 829ms | http |
| 101.47.73.135:3128 | ✓ 853ms | 否 | ✓ 807ms | ✓ 1276ms | ✓ 1444ms | http |
| 222.228.171.92:8080 | ✓ 1630ms | ✓ 1626ms | 否 | ✓ 930ms | ✓ 1685ms | http |
| 103.84.95.54:7890 | ✓ 1185ms | 否 | ✓ 674ms | 否 | ✓ 848ms | http |
| 137.220.151.110:6005 | ✓ 775ms | ✓ 1871ms | ✓ 887ms | 否 | ✓ 929ms | http |
| 120.92.212.16:8890 | ✓ 975ms | ✓ 1193ms | ✓ 941ms | ✓ 1430ms | ✓ 1133ms | http |
| 1.225.116.115:1080 | ✓ 1423ms | ✓ 1085ms | ✓ 853ms | 否 | 否 | http |
| 38.145.220.38:8445 | ✓ 432ms | ✓ 1459ms | ✓ 103ms | ✓ 683ms | ✓ 802ms | http |
| 45.136.130.181:8446 | ✓ 777ms | ✓ 862ms | ✓ 597ms | ✓ 746ms | ✓ 573ms | http |
| 45.136.130.172:8453 | ✓ 434ms | ✓ 752ms | ✓ 696ms | ✓ 1447ms | ✓ 593ms | http |
| 38.145.218.218:8446 | ✓ 1489ms | ✓ 1930ms | ✓ 735ms | ✓ 1013ms | ✓ 563ms | http |
| 38.145.220.65:8450 | 否 | 否 | ✓ 1336ms | ✓ 1224ms | ✓ 605ms | http |
| 101.43.127.100:8877 | ✓ 877ms | ✓ 1051ms | ✓ 879ms | ✓ 1097ms | ✓ 792ms | http |
| 147.161.239.240:8800 | ✓ 1296ms | 否 | ✓ 1192ms | ✓ 1775ms | ✓ 1552ms | http |
| 160.250.4.245:1 | ✓ 1626ms | 否 | 否 | ✓ 1224ms | ✓ 980ms | http |
| 181.78.44.63:999 | ✓ 1000ms | ✓ 1697ms | ✓ 1411ms | ✓ 1640ms | ✓ 1356ms | http |
| 181.41.201.85:3128 | ✓ 738ms | 否 | ✓ 732ms | 否 | ✓ 1741ms | http |
| 106.75.15.167:7890 | ✓ 980ms | ✓ 1672ms | ✓ 1377ms | ✓ 1484ms | ✓ 1155ms | http |
| 45.136.131.54:8448 | ✓ 151ms | ✓ 1343ms | ✓ 957ms | ✓ 805ms | ✓ 528ms | http |
| 142.171.224.229:7890 | ✓ 1430ms | ✓ 678ms | ✓ 847ms | ✓ 1024ms | 否 | http |
| 137.220.150.152:6005 | ✓ 1083ms | 否 | ✓ 930ms | 否 | ✓ 1444ms | http |
| 93.43.251.159:80 | ✓ 1747ms | 否 | ✓ 1745ms | 否 | ✓ 1994ms | http |
| 158.160.215.167:8126 | ✓ 1251ms | 否 | ✓ 1158ms | 否 | ✓ 1616ms | http |
| 158.160.215.167:8124 | ✓ 1263ms | ✓ 1959ms | ✓ 1200ms | 否 | ✓ 1635ms | http |
| 45.136.131.28:8444 | ✓ 1606ms | ✓ 919ms | ✓ 185ms | ✓ 738ms | 否 | http |
| 38.34.179.178:8444 | 否 | ✓ 1136ms | ✓ 1110ms | 否 | ✓ 704ms | http |
| 38.145.208.206:8448 | ✓ 884ms | ✓ 682ms | ✓ 392ms | ✓ 1167ms | 否 | http |
| 62.113.119.14:8080 | ✓ 818ms | 否 | ✓ 791ms | ✓ 1640ms | ✓ 1219ms | http |
| 38.34.179.26:8450 | 否 | ✓ 1730ms | ✓ 1846ms | ✓ 1750ms | ✓ 501ms | http |
| 45.136.130.194:8450 | 否 | ✓ 1224ms | ✓ 1192ms | ✓ 1241ms | ✓ 509ms | http |
| 38.145.208.229:8451 | ✓ 359ms | ✓ 775ms | ✓ 93ms | ✓ 672ms | ✓ 500ms | http |
| 38.34.179.62:8449 | 否 | ✓ 842ms | ✓ 108ms | ✓ 820ms | 否 | http |
| 111.224.193.139:33333 | 否 | ✓ 1172ms | ✓ 1990ms | ✓ 1362ms | 否 | http |
| 38.34.179.54:8446 | ✓ 766ms | ✓ 1053ms | ✓ 1184ms | 否 | 否 | http |
| 193.122.96.242:3128 | ✓ 1125ms | 否 | 否 | ✓ 1302ms | ✓ 1427ms | http |
| 38.34.179.62:8443 | ✓ 405ms | ✓ 648ms | ✓ 270ms | ✓ 660ms | ✓ 555ms | http |
| 43.134.166.79:8888 | ✓ 1539ms | ✓ 1385ms | ✓ 1063ms | ✓ 1331ms | ✓ 1174ms | http |
| 45.174.253.64:999 | ✓ 863ms | ✓ 1858ms | ✓ 1548ms | ✓ 1935ms | ✓ 1233ms | http |
| 38.145.203.108:8446 | ✓ 1036ms | ✓ 1507ms | 否 | ✓ 1798ms | 否 | http |
| 38.34.179.72:8445 | ✓ 593ms | ✓ 689ms | ✓ 423ms | ✓ 717ms | ✓ 648ms | http |
| 45.136.130.247:8450 | ✓ 594ms | ✓ 796ms | ✓ 212ms | ✓ 802ms | ✓ 1029ms | http |
| 38.34.179.193:8448 | ✓ 594ms | ✓ 857ms | ✓ 462ms | ✓ 1104ms | ✓ 637ms | http |
| 38.34.179.181:8448 | ✓ 591ms | ✓ 912ms | ✓ 450ms | ✓ 1094ms | ✓ 648ms | http |
| 38.34.179.185:8451 | ✓ 591ms | ✓ 655ms | ✓ 353ms | ✓ 716ms | ✓ 915ms | http |
| 38.145.208.172:8448 | ✓ 593ms | ✓ 1363ms | ✓ 254ms | ✓ 977ms | ✓ 1074ms | http |
| 38.145.218.234:8445 | ✓ 590ms | ✓ 804ms | ✓ 1406ms | ✓ 1363ms | ✓ 543ms | http |
| 45.136.130.187:8452 | ✓ 588ms | ✓ 1262ms | ✓ 95ms | ✓ 715ms | ✓ 557ms | http |
| 38.145.220.65:8446 | ✓ 1527ms | ✓ 851ms | ✓ 118ms | ✓ 816ms | ✓ 1443ms | http |
| 38.145.220.168:8446 | ✓ 807ms | ✓ 1042ms | ✓ 333ms | 否 | ✓ 1552ms | http |
| 38.34.178.7:8446 | ✓ 1526ms | ✓ 1105ms | ✓ 101ms | ✓ 712ms | ✓ 1031ms | http |
| 38.34.178.193:8449 | ✓ 1287ms | ✓ 1389ms | ✓ 167ms | ✓ 732ms | ✓ 793ms | http |
| 45.136.131.37:8445 | ✓ 1813ms | ✓ 634ms | ✓ 684ms | ✓ 1468ms | ✓ 1745ms | http |
| 45.136.130.171:8448 | 否 | 否 | ✓ 443ms | ✓ 654ms | ✓ 581ms | http |
| 45.136.130.250:8449 | ✓ 1288ms | 否 | ✓ 118ms | ✓ 837ms | ✓ 998ms | http |
| 38.145.218.163:8448 | ✓ 600ms | ✓ 1399ms | ✓ 1152ms | ✓ 687ms | ✓ 728ms | http |
| 45.136.130.168:8448 | 否 | 否 | ✓ 446ms | ✓ 660ms | ✓ 602ms | http |
| 45.136.130.166:8448 | 否 | 否 | ✓ 457ms | ✓ 659ms | ✓ 621ms | http |
| 38.145.203.135:8444 | ✓ 587ms | ✓ 716ms | ✓ 917ms | 否 | ✓ 651ms | http |
| 38.145.208.243:8445 | ✓ 587ms | ✓ 623ms | ✓ 1214ms | 否 | ✓ 563ms | http |
| 45.136.130.167:8446 | ✓ 1526ms | 否 | ✓ 857ms | ✓ 655ms | ✓ 581ms | http |
| 38.34.179.38:8451 | ✓ 1768ms | ✓ 1292ms | 否 | ✓ 1091ms | ✓ 516ms | http |
| 38.34.179.37:8451 | ✓ 1771ms | ✓ 1335ms | 否 | ✓ 1073ms | ✓ 514ms | http |
| 38.34.179.70:8443 | ✓ 98ms | ✓ 605ms | ✓ 81ms | ✓ 683ms | ✓ 505ms | http |
| 38.34.179.65:8443 | ✓ 94ms | ✓ 649ms | ✓ 90ms | ✓ 656ms | ✓ 522ms | http |
| 38.34.179.69:8443 | ✓ 93ms | ✓ 617ms | ✓ 101ms | ✓ 681ms | ✓ 507ms | http |
| 38.34.179.40:8446 | ✓ 1430ms | 否 | ✓ 87ms | ✓ 696ms | ✓ 518ms | http |
| 38.34.179.64:8443 | ✓ 359ms | ✓ 602ms | ✓ 90ms | ✓ 674ms | ✓ 507ms | http |
| 38.34.179.67:8443 | ✓ 359ms | ✓ 599ms | ✓ 88ms | ✓ 678ms | ✓ 517ms | http |
| 38.34.179.186:8444 | ✓ 801ms | ✓ 622ms | ✓ 89ms | ✓ 695ms | ✓ 513ms | http |
| 38.34.183.234:8450 | ✓ 1653ms | 否 | ✓ 96ms | ✓ 667ms | ✓ 525ms | http |
| 38.145.208.208:8447 | 否 | 否 | ✓ 302ms | ✓ 707ms | ✓ 550ms | http |
| 38.145.220.33:8448 | 否 | 否 | ✓ 92ms | ✓ 797ms | ✓ 1179ms | http |
| 8.212.130.232:8080 | 否 | ✓ 1739ms | 否 | ✓ 1186ms | ✓ 886ms | http |
| 202.141.161.53:30001 | ✓ 1028ms | ✓ 1374ms | ✓ 1127ms | ✓ 1186ms | ✓ 993ms | http |
| 59.46.216.131:30001 | ✓ 991ms | ✓ 1277ms | ✓ 1071ms | 否 | 否 | http |
| 49.156.44.114:8080 | ✓ 1471ms | ✓ 1545ms | ✓ 1506ms | ✓ 1344ms | ✓ 1348ms | http |
| 210.223.44.230:3128 | ✓ 1349ms | ✓ 903ms | ✓ 681ms | ✓ 946ms | ✓ 1099ms | http |
| 38.34.178.7:8452 | ✓ 1075ms | ✓ 1167ms | ✓ 824ms | 否 | ✓ 515ms | http |
| 45.136.131.58:8449 | ✓ 705ms | ✓ 1455ms | 否 | ✓ 1570ms | 否 | http |
| 45.136.131.62:8449 | ✓ 669ms | ✓ 1413ms | 否 | ✓ 1620ms | 否 | http |
| 51.79.207.21:8080 | ✓ 874ms | ✓ 1848ms | ✓ 1125ms | ✓ 1947ms | 否 | http |
| 34.101.184.164:3128 | ✓ 803ms | 否 | ✓ 1696ms | ✓ 1387ms | ✓ 1339ms | http |
| 103.82.23.118:5226 | ✓ 1319ms | 否 | ✓ 1242ms | ✓ 1592ms | ✓ 1769ms | http |

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
