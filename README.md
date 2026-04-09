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

最后更新：2026-04-09 16:23:18 UTC（2026-04-10 00:23:18 UTC+8）

**代理总数：247**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 247 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 247 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 20.78.26.206:8561 | ✓ 449ms | ✓ 898ms | ✓ 457ms | ✓ 740ms | ✓ 633ms | http |
| 20.78.118.91:8561 | ✓ 454ms | ✓ 1201ms | ✓ 439ms | ✓ 787ms | ✓ 576ms | http |
| 1.231.81.166:3128 | ✓ 722ms | ✓ 1352ms | ✓ 1358ms | ✓ 1031ms | ✓ 789ms | http |
| 147.161.210.140:8800 | ✓ 710ms | 否 | ✓ 710ms | ✓ 1114ms | ✓ 1142ms | http |
| 35.225.22.61:80 | ✓ 616ms | 否 | ✓ 1255ms | ✓ 1479ms | ✓ 951ms | http |
| 34.96.238.40:8080 | ✓ 1053ms | 否 | ✓ 1128ms | ✓ 1716ms | 否 | http |
| 154.40.137.209:55965 | ✓ 715ms | 否 | ✓ 473ms | ✓ 1937ms | 否 | http |
| 38.92.10.152:57579 | ✓ 620ms | 否 | 否 | ✓ 1026ms | ✓ 956ms | http |
| 161.35.70.36:8888 | ✓ 867ms | 否 | ✓ 973ms | 否 | ✓ 1381ms | http |
| 113.160.132.26:8080 | ✓ 1856ms | 否 | ✓ 1246ms | ✓ 1185ms | ✓ 947ms | http |
| 212.58.132.5:8888 | ✓ 1228ms | 否 | ✓ 1690ms | ✓ 1586ms | ✓ 1275ms | http |
| 38.34.179.79:8449 | ✓ 180ms | ✓ 735ms | ✓ 307ms | ✓ 695ms | ✓ 539ms | http |
| 45.136.131.28:8449 | ✓ 237ms | ✓ 915ms | ✓ 121ms | ✓ 926ms | ✓ 562ms | http |
| 38.145.218.162:8448 | ✓ 184ms | ✓ 667ms | ✓ 706ms | ✓ 929ms | ✓ 536ms | http |
| 38.34.179.64:8451 | ✓ 294ms | ✓ 1422ms | ✓ 181ms | ✓ 749ms | ✓ 591ms | http |
| 38.145.203.39:8445 | ✓ 210ms | ✓ 690ms | ✓ 674ms | ✓ 706ms | ✓ 540ms | http |
| 38.145.218.160:8448 | ✓ 142ms | 否 | ✓ 87ms | ✓ 694ms | ✓ 519ms | http |
| 38.145.208.224:8445 | ✓ 159ms | ✓ 1790ms | ✓ 399ms | ✓ 909ms | ✓ 740ms | http |
| 38.92.10.139:33985 | ✓ 561ms | ✓ 961ms | ✓ 888ms | 否 | 否 | http |
| 38.244.54.190:31168 | ✓ 843ms | ✓ 725ms | ✓ 796ms | ✓ 1270ms | ✓ 693ms | http |
| 20.210.39.153:8561 | ✓ 648ms | 否 | ✓ 454ms | ✓ 813ms | ✓ 629ms | http |
| 38.145.208.211:8453 | ✓ 724ms | ✓ 1317ms | ✓ 549ms | ✓ 1014ms | ✓ 885ms | http |
| 104.234.0.145:55554 | ✓ 620ms | ✓ 1230ms | ✓ 710ms | ✓ 1227ms | ✓ 761ms | http |
| 38.34.179.194:8451 | 否 | 否 | ✓ 1100ms | ✓ 725ms | ✓ 911ms | http |
| 8.209.239.31:30000 | ✓ 459ms | 否 | 否 | ✓ 1595ms | ✓ 553ms | http |
| 167.103.34.108:8800 | ✓ 1988ms | 否 | ✓ 1094ms | ✓ 1323ms | ✓ 1222ms | http |
| 45.88.0.113:3128 | ✓ 739ms | 否 | ✓ 1232ms | 否 | ✓ 1237ms | http |
| 103.3.246.71:3128 | ✓ 1294ms | 否 | ✓ 1700ms | ✓ 1493ms | ✓ 961ms | http |
| 45.88.0.111:3128 | ✓ 702ms | 否 | ✓ 1294ms | 否 | ✓ 1145ms | http |
| 45.88.0.115:3128 | ✓ 706ms | 否 | ✓ 1213ms | 否 | ✓ 1163ms | http |
| 45.88.0.98:3128 | ✓ 715ms | 否 | ✓ 1213ms | 否 | ✓ 1258ms | http |
| 59.46.216.131:30001 | ✓ 917ms | ✓ 1305ms | ✓ 1324ms | 否 | ✓ 1041ms | http |
| 45.88.0.117:3128 | ✓ 744ms | 否 | 否 | ✓ 1550ms | ✓ 1329ms | http |
| 45.88.0.114:3128 | ✓ 691ms | 否 | ✓ 1671ms | 否 | ✓ 1214ms | http |
| 16.78.119.130:443 | 否 | ✓ 1588ms | 否 | ✓ 1657ms | ✓ 1477ms | http |
| 45.167.124.52:8080 | ✓ 1392ms | 否 | ✓ 713ms | ✓ 1742ms | 否 | http |
| 167.103.144.127:8800 | ✓ 779ms | 否 | ✓ 732ms | ✓ 1197ms | ✓ 1076ms | http |
| 120.132.52.172:8888 | 否 | 否 | ✓ 1809ms | ✓ 1222ms | ✓ 1137ms | http |
| 185.191.236.162:3128 | ✓ 724ms | 否 | ✓ 782ms | ✓ 1779ms | ✓ 1194ms | http |
| 162.240.154.26:3128 | ✓ 983ms | ✓ 1898ms | 否 | ✓ 1000ms | ✓ 1961ms | http |
| 115.231.181.40:8128 | ✓ 963ms | ✓ 1721ms | ✓ 882ms | 否 | ✓ 1342ms | http |
| 93.77.179.177:8888 | ✓ 831ms | ✓ 1996ms | ✓ 860ms | 否 | ✓ 1637ms | http |
| 167.103.31.122:8800 | ✓ 1304ms | 否 | ✓ 1370ms | ✓ 1637ms | ✓ 1469ms | http |
| 45.12.151.226:2829 | ✓ 909ms | 否 | ✓ 1841ms | ✓ 1804ms | 否 | http |
| 91.238.105.64:2024 | ✓ 1233ms | 否 | ✓ 1766ms | 否 | ✓ 1985ms | http |
| 137.59.47.73:3128 | ✓ 1400ms | 否 | ✓ 1396ms | 否 | ✓ 1138ms | http |
| 152.32.132.190:7890 | ✓ 1568ms | 否 | ✓ 648ms | ✓ 1429ms | 否 | http |
| 45.136.131.42:8444 | 否 | ✓ 1697ms | ✓ 168ms | 否 | ✓ 1278ms | http |
| 58.147.186.226:8097 | ✓ 1957ms | 否 | 否 | ✓ 1446ms | ✓ 1692ms | http |
| 210.223.44.230:3128 | ✓ 1424ms | ✓ 1266ms | ✓ 717ms | ✓ 938ms | ✓ 723ms | http |
| 38.34.179.179:8448 | ✓ 769ms | 否 | ✓ 302ms | ✓ 674ms | ✓ 630ms | http |
| 38.145.208.174:8444 | ✓ 787ms | ✓ 1507ms | ✓ 1073ms | ✓ 877ms | ✓ 961ms | http |
| 106.10.55.212:1121 | ✓ 1422ms | 否 | ✓ 1198ms | ✓ 1184ms | ✓ 939ms | http |
| 45.136.131.68:8451 | ✓ 1629ms | 否 | ✓ 888ms | ✓ 1971ms | ✓ 736ms | http |
| 147.161.239.240:8800 | ✓ 905ms | ✓ 1852ms | ✓ 1454ms | ✓ 1693ms | ✓ 1794ms | http |
| 5.104.87.17:8051 | ✓ 1251ms | 否 | ✓ 1755ms | ✓ 1523ms | 否 | http |
| 101.43.127.100:8877 | 否 | ✓ 1016ms | ✓ 819ms | ✓ 1058ms | ✓ 1856ms | http |
| 45.88.0.99:3128 | ✓ 731ms | 否 | ✓ 1944ms | 否 | ✓ 1578ms | http |
| 213.220.62.62:3128 | ✓ 725ms | 否 | ✓ 1953ms | 否 | ✓ 1582ms | http |
| 45.88.0.116:3128 | ✓ 724ms | 否 | ✓ 1943ms | 否 | ✓ 1591ms | http |
| 213.220.62.63:3128 | ✓ 745ms | 否 | ✓ 1928ms | 否 | ✓ 1558ms | http |
| 202.47.185.178:8080 | ✓ 1915ms | 否 | ✓ 1498ms | ✓ 1306ms | ✓ 1249ms | http |
| 181.78.44.63:999 | 否 | ✓ 1608ms | ✓ 1482ms | ✓ 1840ms | ✓ 1571ms | http |
| 45.167.125.21:999 | ✓ 1428ms | 否 | ✓ 1487ms | 否 | ✓ 1781ms | http |
| 177.234.217.88:999 | ✓ 1537ms | 否 | 否 | ✓ 1965ms | ✓ 1721ms | http |
| 168.110.52.228:3128 | ✓ 609ms | ✓ 980ms | ✓ 508ms | ✓ 722ms | ✓ 583ms | http |
| 38.145.218.101:8444 | ✓ 160ms | ✓ 659ms | ✓ 889ms | ✓ 1214ms | ✓ 1628ms | http |
| 38.34.183.222:8453 | ✓ 1820ms | 否 | ✓ 1076ms | ✓ 988ms | ✓ 513ms | http |
| 38.145.208.238:8447 | ✓ 176ms | ✓ 670ms | ✓ 765ms | ✓ 707ms | ✓ 734ms | http |
| 38.145.208.235:8447 | ✓ 192ms | ✓ 693ms | ✓ 767ms | ✓ 718ms | ✓ 703ms | http |
| 91.238.104.171:2023 | ✓ 1167ms | 否 | ✓ 1808ms | ✓ 1781ms | ✓ 1364ms | http |
| 38.34.179.150:8449 | ✓ 1768ms | ✓ 1424ms | 否 | ✓ 1661ms | ✓ 1425ms | http |
| 38.145.208.227:8447 | ✓ 180ms | ✓ 1655ms | ✓ 165ms | ✓ 806ms | ✓ 1558ms | http |
| 38.145.208.237:8447 | ✓ 197ms | 否 | ✓ 215ms | ✓ 947ms | ✓ 1494ms | http |
| 38.145.208.204:8446 | ✓ 765ms | ✓ 1338ms | ✓ 191ms | ✓ 923ms | ✓ 653ms | http |
| 45.136.130.183:8447 | ✓ 1277ms | ✓ 658ms | ✓ 145ms | ✓ 752ms | ✓ 1062ms | http |
| 45.136.130.177:8447 | ✓ 1592ms | ✓ 634ms | ✓ 170ms | ✓ 970ms | ✓ 1428ms | http |
| 45.136.130.193:8447 | ✓ 788ms | 否 | ✓ 207ms | ✓ 1805ms | ✓ 857ms | http |
| 142.171.95.105:3128 | ✓ 1568ms | ✓ 1844ms | 否 | ✓ 696ms | ✓ 685ms | http |
| 171.232.59.181:4004 | ✓ 1438ms | 否 | ✓ 941ms | ✓ 1183ms | ✓ 1111ms | http |
| 120.92.212.16:7890 | ✓ 929ms | ✓ 1143ms | ✓ 1011ms | 否 | ✓ 1214ms | http |
| 222.228.171.92:8080 | ✓ 1917ms | 否 | ✓ 1443ms | 否 | ✓ 820ms | http |
| 38.145.208.226:8448 | ✓ 738ms | ✓ 1201ms | ✓ 960ms | 否 | ✓ 1273ms | http |
| 172.104.63.237:3128 | ✓ 1543ms | 否 | ✓ 837ms | ✓ 1530ms | ✓ 1360ms | http |
| 120.92.212.16:8890 | ✓ 1178ms | 否 | ✓ 1761ms | 否 | ✓ 960ms | http |
| 116.254.118.180:80 | ✓ 1899ms | 否 | 否 | ✓ 1215ms | ✓ 1871ms | http |
| 45.136.131.40:8444 | ✓ 1881ms | ✓ 974ms | ✓ 926ms | ✓ 1909ms | ✓ 972ms | http |
| 38.34.179.24:8447 | ✓ 1855ms | 否 | ✓ 1547ms | 否 | ✓ 1520ms | http |
| 85.239.59.252:7890 | ✓ 1997ms | 否 | ✓ 1806ms | 否 | ✓ 1612ms | http |
| 160.25.46.58:1111 | ✓ 1683ms | 否 | 否 | ✓ 1810ms | ✓ 1707ms | http |
| 5.104.87.17:8050 | ✓ 1559ms | 否 | ✓ 984ms | ✓ 1187ms | ✓ 800ms | http |
| 167.103.115.102:8800 | ✓ 1418ms | 否 | ✓ 897ms | ✓ 1264ms | ✓ 967ms | http |
| 150.107.140.238:3128 | ✓ 1618ms | 否 | 否 | ✓ 1575ms | ✓ 1550ms | http |
| 20.210.76.178:8561 | ✓ 1174ms | ✓ 1855ms | ✓ 1816ms | 否 | 否 | http |
| 20.210.76.104:8561 | ✓ 1178ms | ✓ 1918ms | ✓ 1806ms | 否 | 否 | http |
| 38.34.179.83:8448 | ✓ 824ms | ✓ 689ms | ✓ 521ms | 否 | 否 | http |
| 38.34.179.85:8444 | ✓ 1110ms | ✓ 803ms | ✓ 129ms | 否 | 否 | http |
| 38.145.208.241:8447 | ✓ 821ms | ✓ 999ms | ✓ 91ms | ✓ 664ms | ✓ 518ms | http |
| 38.145.220.81:8445 | ✓ 800ms | ✓ 691ms | 否 | ✓ 729ms | ✓ 547ms | http |
| 38.34.179.94:8444 | ✓ 1110ms | ✓ 812ms | ✓ 405ms | ✓ 735ms | ✓ 1151ms | http |

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
