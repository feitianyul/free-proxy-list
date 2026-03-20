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

最后更新：2026-03-20 10:40:18 UTC（2026-03-20 18:40:18 UTC+8）

**代理总数：95**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 95 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 95 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 178.156.187.185:10001 | ✓ 331ms | ✓ 1355ms | ✓ 757ms | ✓ 1630ms | ✓ 1237ms | http |
| 219.117.204.211:7799 | ✓ 1472ms | 否 | ✓ 528ms | ✓ 1104ms | ✓ 853ms | http |
| 147.161.210.140:8800 | ✓ 1470ms | ✓ 1992ms | ✓ 823ms | ✓ 917ms | ✓ 1003ms | http |
| 1.231.81.166:3128 | ✓ 1505ms | ✓ 983ms | ✓ 1319ms | ✓ 1108ms | ✓ 966ms | http |
| 147.161.239.240:8800 | ✓ 1197ms | ✓ 1974ms | ✓ 1543ms | ✓ 1664ms | ✓ 1657ms | http |
| 113.160.132.26:8080 | ✓ 1986ms | ✓ 1309ms | ✓ 1652ms | ✓ 1241ms | ✓ 1029ms | http |
| 103.166.185.54:3128 | ✓ 1979ms | 否 | ✓ 1547ms | ✓ 1196ms | ✓ 948ms | http |
| 38.34.179.162:8451 | ✓ 162ms | ✓ 1188ms | ✓ 1099ms | ✓ 792ms | ✓ 531ms | http |
| 167.103.34.108:8800 | ✓ 1140ms | 否 | ✓ 1244ms | ✓ 1403ms | ✓ 1252ms | http |
| 137.220.150.22:6005 | 否 | 否 | ✓ 1247ms | ✓ 1609ms | ✓ 1402ms | http |
| 162.240.154.26:3128 | ✓ 522ms | ✓ 1858ms | 否 | ✓ 1420ms | 否 | http |
| 103.84.95.54:7890 | ✓ 1016ms | 否 | ✓ 797ms | ✓ 828ms | ✓ 1306ms | http |
| 154.64.243.50:7890 | 否 | ✓ 1415ms | ✓ 128ms | ✓ 923ms | ✓ 682ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1162ms | ✓ 920ms | ✓ 1180ms | 否 | http |
| 183.249.5.117:22222 | ✓ 813ms | ✓ 1429ms | ✓ 1015ms | ✓ 1541ms | ✓ 1468ms | http |
| 85.208.108.43:2094 | 否 | 否 | ✓ 893ms | ✓ 1201ms | ✓ 964ms | http |
| 137.220.151.110:6005 | 否 | 否 | ✓ 1198ms | ✓ 1081ms | ✓ 843ms | http |
| 167.103.31.122:8800 | ✓ 1422ms | 否 | ✓ 1405ms | 否 | ✓ 1534ms | http |
| 106.75.15.167:7890 | 否 | ✓ 1122ms | ✓ 894ms | 否 | ✓ 1617ms | http |
| 111.201.98.211:7890 | ✓ 1084ms | ✓ 1106ms | ✓ 1667ms | ✓ 1154ms | ✓ 930ms | http |
| 139.99.238.95:8080 | ✓ 1390ms | 否 | ✓ 713ms | ✓ 1227ms | ✓ 956ms | http |
| 158.160.215.167:8124 | ✓ 1424ms | 否 | ✓ 816ms | 否 | ✓ 1643ms | http |
| 102.134.48.240:6005 | ✓ 1530ms | ✓ 1307ms | ✓ 1292ms | ✓ 1259ms | ✓ 933ms | http |
| 38.145.218.230:8453 | 否 | ✓ 1443ms | ✓ 1107ms | ✓ 1860ms | ✓ 671ms | http |
| 101.43.127.100:8877 | ✓ 881ms | ✓ 1076ms | ✓ 1061ms | ✓ 1118ms | ✓ 809ms | http |
| 121.138.61.193:8818 | ✓ 1558ms | ✓ 1639ms | ✓ 1878ms | ✓ 1012ms | ✓ 797ms | http |
| 101.47.73.135:3128 | ✓ 1703ms | 否 | ✓ 1005ms | ✓ 1531ms | ✓ 1068ms | http |
| 210.45.70.16:7895 | ✓ 1155ms | ✓ 1335ms | ✓ 1621ms | ✓ 1637ms | ✓ 1069ms | http |
| 194.67.99.223:1080 | ✓ 1162ms | 否 | ✓ 1848ms | 否 | ✓ 1580ms | http |
| 120.92.212.16:7890 | ✓ 967ms | 否 | ✓ 1399ms | 否 | ✓ 1207ms | http |
| 120.92.212.16:8890 | ✓ 945ms | 否 | 否 | ✓ 1193ms | ✓ 985ms | http |
| 16.78.119.130:443 | 否 | 否 | ✓ 1812ms | ✓ 1960ms | ✓ 1435ms | http |
| 137.220.150.170:6005 | ✓ 827ms | 否 | 否 | ✓ 1264ms | ✓ 1354ms | http |
| 103.139.138.194:3128 | 否 | 否 | ✓ 1072ms | ✓ 1443ms | ✓ 1056ms | http |
| 185.41.152.110:3128 | ✓ 1134ms | ✓ 1715ms | ✓ 909ms | ✓ 1657ms | ✓ 1889ms | http |
| 45.136.131.44:8452 | ✓ 565ms | ✓ 1178ms | ✓ 635ms | ✓ 1146ms | ✓ 634ms | http |
| 38.145.218.102:8444 | ✓ 575ms | ✓ 1294ms | ✓ 509ms | ✓ 692ms | ✓ 1543ms | http |
| 91.238.104.171:2023 | ✓ 886ms | 否 | ✓ 854ms | ✓ 1698ms | ✓ 1295ms | http |
| 194.147.115.50:3128 | ✓ 746ms | 否 | ✓ 934ms | 否 | ✓ 1599ms | http |
| 38.34.179.24:8453 | ✓ 970ms | ✓ 1357ms | ✓ 543ms | ✓ 1696ms | ✓ 1968ms | http |
| 183.249.5.111:22222 | 否 | 否 | ✓ 832ms | ✓ 966ms | ✓ 802ms | http |
| 183.249.5.110:22222 | ✓ 788ms | 否 | 否 | ✓ 1251ms | ✓ 711ms | http |
| 62.113.119.14:8080 | ✓ 810ms | 否 | ✓ 730ms | ✓ 1592ms | ✓ 1194ms | http |
| 198.244.188.127:3128 | ✓ 915ms | 否 | ✓ 1782ms | 否 | ✓ 1755ms | http |
| 174.138.24.77:1080 | ✓ 1572ms | 否 | ✓ 1156ms | ✓ 1545ms | ✓ 1226ms | http |
| 45.136.131.58:8450 | ✓ 652ms | ✓ 1709ms | ✓ 494ms | ✓ 750ms | ✓ 996ms | http |
| 45.136.131.59:8450 | ✓ 383ms | ✓ 808ms | ✓ 1207ms | ✓ 1062ms | 否 | http |
| 137.220.150.152:6005 | ✓ 1488ms | 否 | ✓ 796ms | ✓ 1183ms | ✓ 905ms | http |
| 35.225.22.61:80 | ✓ 428ms | ✓ 1478ms | 否 | ✓ 1461ms | ✓ 972ms | http |
| 38.34.178.154:8445 | ✓ 185ms | ✓ 1128ms | ✓ 1112ms | 否 | ✓ 757ms | http |
| 144.31.253.242:1080 | ✓ 1284ms | 否 | ✓ 618ms | ✓ 1764ms | 否 | http |
| 38.34.179.105:8449 | ✓ 1232ms | ✓ 1313ms | ✓ 314ms | ✓ 691ms | ✓ 543ms | http |
| 91.238.105.64:2024 | ✓ 1721ms | ✓ 1819ms | ✓ 1836ms | 否 | 否 | http |
| 137.220.150.104:6005 | ✓ 788ms | 否 | ✓ 1206ms | ✓ 1340ms | ✓ 931ms | http |
| 133.242.138.34:8100 | ✓ 1487ms | 否 | ✓ 1758ms | ✓ 1819ms | 否 | http |
| 183.249.5.105:22222 | ✓ 894ms | ✓ 1192ms | ✓ 670ms | ✓ 1130ms | ✓ 681ms | http |
| 217.76.245.80:999 | ✓ 1403ms | 否 | ✓ 1227ms | ✓ 1672ms | ✓ 1456ms | http |
| 101.32.244.83:8080 | ✓ 1217ms | 否 | ✓ 932ms | ✓ 1238ms | ✓ 1277ms | http |
| 121.43.196.213:8222 | ✓ 955ms | ✓ 1043ms | ✓ 970ms | ✓ 1087ms | ✓ 845ms | http |
| 121.43.196.210:8222 | ✓ 993ms | ✓ 1092ms | ✓ 876ms | ✓ 1094ms | ✓ 901ms | http |
| 114.55.226.123:10086 | ✓ 1076ms | ✓ 1438ms | ✓ 1101ms | ✓ 1270ms | ✓ 1038ms | http |
| 45.136.198.40:3128 | ✓ 1400ms | 否 | ✓ 826ms | 否 | ✓ 1920ms | http |
| 38.145.218.101:8448 | 否 | ✓ 660ms | ✓ 373ms | ✓ 1004ms | ✓ 1402ms | http |
| 38.145.220.179:8450 | 否 | 否 | ✓ 773ms | ✓ 684ms | ✓ 521ms | http |
| 38.145.220.60:8444 | 否 | ✓ 769ms | ✓ 585ms | ✓ 1879ms | ✓ 512ms | http |
| 38.34.179.63:8451 | 否 | 否 | ✓ 1295ms | ✓ 818ms | ✓ 558ms | http |
| 38.145.203.135:8451 | 否 | 否 | ✓ 383ms | ✓ 1220ms | ✓ 1576ms | http |
| 222.228.171.92:8080 | ✓ 1600ms | ✓ 1703ms | ✓ 549ms | ✓ 1891ms | 否 | http |
| 172.212.68.37:3128 | ✓ 1120ms | 否 | ✓ 857ms | ✓ 1671ms | ✓ 1148ms | http |
| 91.233.223.147:3128 | ✓ 1528ms | 否 | ✓ 1057ms | 否 | ✓ 1676ms | http |
| 149.56.24.51:3128 | ✓ 353ms | ✓ 1494ms | ✓ 445ms | ✓ 1264ms | ✓ 970ms | http |
| 47.77.193.180:1080 | ✓ 676ms | 否 | ✓ 374ms | ✓ 725ms | ✓ 548ms | http |
| 222.184.48.251:22222 | ✓ 1957ms | ✓ 1237ms | ✓ 977ms | 否 | 否 | http |
| 59.46.216.131:30001 | ✓ 1987ms | 否 | ✓ 1913ms | ✓ 1372ms | 否 | http |
| 8.212.172.106:8080 | ✓ 1525ms | 否 | ✓ 1186ms | ✓ 1303ms | 否 | http |
| 222.184.48.252:22222 | ✓ 1540ms | ✓ 1804ms | ✓ 1217ms | 否 | ✓ 959ms | http |
| 45.144.28.81:10808 | ✓ 1157ms | 否 | ✓ 1303ms | ✓ 1841ms | 否 | http |
| 5.102.109.41:999 | 否 | 否 | ✓ 1318ms | ✓ 1281ms | ✓ 1193ms | http |
| 45.168.238.193:8443 | ✓ 1370ms | 否 | ✓ 267ms | 否 | ✓ 934ms | http |
| 185.114.73.2:1080 | ✓ 1130ms | ✓ 1918ms | ✓ 1308ms | ✓ 1660ms | ✓ 1316ms | http |
| 45.93.29.147:6005 | ✓ 1352ms | 否 | ✓ 1553ms | ✓ 1139ms | ✓ 1017ms | http |
| 38.34.179.97:8448 | ✓ 1100ms | ✓ 1646ms | ✓ 1709ms | ✓ 1111ms | 否 | http |
| 61.52.131.172:8443 | ✓ 848ms | ✓ 1143ms | ✓ 937ms | ✓ 1182ms | ✓ 917ms | http |
| 114.237.77.231:1080 | ✓ 953ms | ✓ 1211ms | ✓ 908ms | ✓ 1886ms | 否 | http |
| 38.34.179.20:8445 | ✓ 734ms | ✓ 689ms | ✓ 1724ms | ✓ 1000ms | ✓ 562ms | http |
| 38.34.179.16:8451 | ✓ 733ms | ✓ 731ms | ✓ 1683ms | ✓ 791ms | ✓ 639ms | http |
| 45.136.130.173:8448 | ✓ 1131ms | ✓ 1705ms | ✓ 869ms | ✓ 747ms | ✓ 1144ms | http |
| 38.145.208.172:8448 | ✓ 1131ms | ✓ 1186ms | ✓ 724ms | ✓ 754ms | ✓ 1053ms | http |
| 45.93.31.93:6005 | 否 | ✓ 1849ms | 否 | ✓ 1562ms | ✓ 1354ms | http |
| 106.117.208.101:7890 | ✓ 1631ms | ✓ 1227ms | 否 | 否 | ✓ 1064ms | http |
| 38.145.220.11:8445 | ✓ 467ms | ✓ 674ms | ✓ 998ms | ✓ 848ms | ✓ 565ms | http |
| 38.34.178.155:8448 | ✓ 394ms | ✓ 723ms | ✓ 188ms | ✓ 904ms | ✓ 1039ms | http |
| 45.136.130.171:8445 | ✓ 608ms | 否 | ✓ 793ms | ✓ 688ms | ✓ 566ms | http |
| 38.145.220.33:8448 | ✓ 462ms | 否 | ✓ 378ms | ✓ 1235ms | ✓ 725ms | http |
| 45.136.131.66:8445 | ✓ 284ms | ✓ 1004ms | ✓ 1617ms | ✓ 999ms | ✓ 1003ms | http |

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
