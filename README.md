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

最后更新：2026-05-05 17:18:01 UTC（2026-05-06 01:18:01 UTC+8）

**代理总数：93**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 93 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 93 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 47.85.51.197:1080 | ✓ 247ms | ✓ 976ms | 否 | ✓ 1293ms | ✓ 695ms | http |
| 34.71.229.255:3128 | ✓ 744ms | ✓ 1590ms | ✓ 1149ms | ✓ 1503ms | ✓ 958ms | http |
| 223.84.151.86:30005 | ✓ 1630ms | ✓ 1302ms | ✓ 1615ms | 否 | ✓ 1656ms | http |
| 152.32.132.190:7890 | ✓ 886ms | 否 | ✓ 1134ms | ✓ 1401ms | ✓ 1869ms | http |
| 148.230.4.241:999 | ✓ 1028ms | ✓ 1781ms | ✓ 807ms | ✓ 1703ms | ✓ 1505ms | http |
| 181.119.97.24:999 | ✓ 1217ms | 否 | ✓ 1795ms | 否 | ✓ 1500ms | http |
| 43.99.54.236:5555 | ✓ 850ms | ✓ 1179ms | ✓ 823ms | ✓ 1068ms | ✓ 841ms | http |
| 80.92.204.47:1081 | ✓ 1031ms | ✓ 1893ms | ✓ 706ms | 否 | ✓ 1790ms | http |
| 91.238.104.171:2023 | ✓ 856ms | 否 | ✓ 1492ms | 否 | ✓ 1569ms | http |
| 84.47.150.125:1080 | ✓ 571ms | 否 | ✓ 1514ms | 否 | ✓ 1708ms | http |
| 14.143.222.113:57788 | ✓ 1816ms | 否 | ✓ 951ms | ✓ 1395ms | 否 | http |
| 45.125.67.37:8443 | ✓ 1087ms | 否 | ✓ 1812ms | 否 | ✓ 1535ms | http |
| 113.160.132.26:8080 | ✓ 1966ms | ✓ 1569ms | ✓ 1580ms | 否 | ✓ 1176ms | http |
| 103.156.74.163:8080 | ✓ 1851ms | 否 | ✓ 1789ms | ✓ 1822ms | ✓ 1821ms | http |
| 149.56.24.51:3128 | ✓ 452ms | 否 | ✓ 1568ms | 否 | ✓ 1213ms | http |
| 8.219.97.248:80 | ✓ 1300ms | 否 | ✓ 1159ms | ✓ 1814ms | 否 | http |
| 183.98.143.134:8097 | 否 | 否 | ✓ 1471ms | ✓ 1302ms | ✓ 1048ms | http |
| 206.206.126.177:2412 | ✓ 1720ms | 否 | ✓ 1846ms | ✓ 1970ms | 否 | http |
| 45.173.12.140:1994 | 否 | 否 | ✓ 1157ms | ✓ 1805ms | ✓ 1517ms | http |
| 38.7.18.188:999 | ✓ 449ms | 否 | ✓ 366ms | ✓ 1120ms | ✓ 919ms | http |
| 116.80.49.66:3172 | ✓ 1728ms | 否 | ✓ 1740ms | 否 | ✓ 1815ms | http |
| 103.171.241.6:8080 | ✓ 1903ms | 否 | ✓ 1506ms | ✓ 1832ms | ✓ 1583ms | http |
| 120.92.212.16:8890 | ✓ 1080ms | ✓ 1623ms | 否 | 否 | ✓ 1213ms | http |
| 120.92.212.16:7890 | ✓ 1205ms | 否 | ✓ 1079ms | ✓ 1836ms | ✓ 1701ms | http |
| 45.225.204.11:999 | ✓ 916ms | 否 | ✓ 630ms | ✓ 1962ms | ✓ 1623ms | http |
| 20.78.118.91:8561 | 否 | ✓ 1303ms | ✓ 885ms | ✓ 1177ms | ✓ 1009ms | http |
| 20.210.39.153:8561 | 否 | ✓ 1263ms | ✓ 921ms | ✓ 1183ms | ✓ 991ms | http |
| 20.78.26.206:8561 | 否 | 否 | ✓ 759ms | ✓ 1115ms | ✓ 884ms | http |
| 183.98.143.134:8084 | 否 | ✓ 1471ms | ✓ 1307ms | ✓ 1434ms | ✓ 1096ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1304ms | ✓ 1392ms | ✓ 1446ms | ✓ 1410ms | http |
| 158.160.215.167:8127 | ✓ 1322ms | 否 | ✓ 1878ms | 否 | ✓ 1679ms | http |
| 45.153.231.229:8080 | ✓ 680ms | 否 | ✓ 1289ms | ✓ 1968ms | 否 | http |
| 20.164.75.153:8080 | ✓ 1141ms | 否 | ✓ 1844ms | 否 | ✓ 1999ms | http |
| 104.129.194.38:8800 | ✓ 371ms | 否 | 否 | ✓ 1910ms | ✓ 1036ms | http |
| 152.70.91.193:40000 | ✓ 1566ms | 否 | ✓ 1827ms | ✓ 1882ms | 否 | http |
| 103.157.200.126:3128 | ✓ 1744ms | 否 | ✓ 1608ms | 否 | ✓ 1952ms | http |
| 217.182.195.221:30003 | ✓ 723ms | 否 | ✓ 1811ms | 否 | ✓ 1756ms | http |
| 210.223.44.230:3128 | ✓ 1742ms | ✓ 1334ms | 否 | 否 | ✓ 1125ms | http |
| 59.46.216.131:30001 | ✓ 1122ms | ✓ 1667ms | 否 | ✓ 1551ms | ✓ 1233ms | http |
| 103.209.36.58:8080 | 否 | 否 | ✓ 1870ms | ✓ 1529ms | ✓ 1409ms | http |
| 116.118.48.147:3128 | ✓ 1595ms | 否 | ✓ 1350ms | ✓ 1413ms | ✓ 1128ms | http |
| 8.213.34.109:3128 | ✓ 1292ms | 否 | ✓ 1141ms | ✓ 1659ms | ✓ 1319ms | http |
| 43.133.44.89:8888 | ✓ 934ms | 否 | ✓ 905ms | 否 | ✓ 963ms | http |
| 150.107.140.238:3128 | ✓ 1299ms | 否 | ✓ 1167ms | ✓ 1613ms | 否 | http |
| 5.63.111.238:8080 | ✓ 1374ms | 否 | ✓ 1214ms | ✓ 1828ms | ✓ 1611ms | http |
| 64.188.77.221:3128 | ✓ 552ms | ✓ 1809ms | ✓ 557ms | ✓ 1902ms | 否 | http |
| 20.27.11.248:8561 | ✓ 653ms | 否 | ✓ 926ms | ✓ 1079ms | ✓ 794ms | http |
| 20.27.15.111:8561 | ✓ 707ms | 否 | ✓ 893ms | ✓ 1072ms | ✓ 783ms | http |
| 20.27.13.35:8561 | ✓ 651ms | ✓ 1005ms | ✓ 656ms | ✓ 1007ms | ✓ 768ms | http |
| 212.58.132.5:8888 | ✓ 1110ms | 否 | ✓ 1225ms | ✓ 1503ms | ✓ 1216ms | http |
| 20.27.14.220:8561 | ✓ 1654ms | ✓ 1422ms | ✓ 693ms | ✓ 1006ms | ✓ 791ms | http |
| 121.230.8.22:1080 | ✓ 1379ms | 否 | ✓ 1185ms | ✓ 1768ms | ✓ 1853ms | http |
| 45.186.6.104:3128 | ✓ 1446ms | ✓ 1779ms | ✓ 1576ms | 否 | 否 | http |
| 20.210.76.178:8561 | 否 | ✓ 1175ms | ✓ 1015ms | ✓ 1033ms | ✓ 892ms | http |
| 20.210.76.104:8561 | ✓ 1529ms | 否 | ✓ 673ms | 否 | ✓ 897ms | http |
| 190.12.150.244:999 | ✓ 737ms | 否 | ✓ 1151ms | ✓ 1673ms | 否 | http |
| 115.231.181.40:8128 | 否 | ✓ 1742ms | ✓ 1678ms | 否 | ✓ 1197ms | http |
| 152.42.177.32:8888 | ✓ 1181ms | 否 | ✓ 1413ms | ✓ 1549ms | ✓ 1570ms | http |
| 113.192.12.24:8080 | ✓ 1616ms | 否 | 否 | ✓ 1716ms | ✓ 1645ms | http |
| 150.249.255.91:3128 | 否 | ✓ 1174ms | 否 | ✓ 1084ms | ✓ 1892ms | http |
| 101.6.65.112:10080 | ✓ 1150ms | ✓ 1485ms | ✓ 1253ms | ✓ 1485ms | ✓ 1491ms | http |
| 3.101.133.120:80 | ✓ 723ms | 否 | ✓ 1452ms | ✓ 1484ms | 否 | http |
| 20.205.16.149:3128 | ✓ 962ms | ✓ 1510ms | ✓ 1060ms | ✓ 1297ms | ✓ 921ms | http |
| 107.173.42.121:7890 | 否 | ✓ 928ms | ✓ 345ms | ✓ 964ms | 否 | http |
| 20.2.83.243:3128 | ✓ 818ms | ✓ 1700ms | ✓ 945ms | ✓ 1143ms | ✓ 903ms | http |
| 199.68.217.2:3128 | ✓ 586ms | ✓ 1174ms | ✓ 1350ms | ✓ 986ms | 否 | http |
| 38.188.247.12:999 | ✓ 1722ms | 否 | 否 | ✓ 1345ms | ✓ 1180ms | http |
| 160.238.65.9:3128 | ✓ 1825ms | 否 | ✓ 1645ms | ✓ 1935ms | ✓ 1343ms | http |
| 134.209.153.66:3128 | ✓ 1542ms | 否 | ✓ 977ms | ✓ 1650ms | ✓ 1467ms | http |
| 216.106.179.216:49152 | 否 | 否 | ✓ 768ms | ✓ 1013ms | ✓ 734ms | http |
| 91.233.223.147:3128 | ✓ 829ms | ✓ 1883ms | ✓ 807ms | 否 | ✓ 1541ms | http |
| 218.108.131.186:17890 | ✓ 990ms | ✓ 1272ms | ✓ 1061ms | ✓ 1312ms | ✓ 1075ms | http |
| 45.129.141.143:3128 | ✓ 1419ms | 否 | ✓ 1702ms | ✓ 1744ms | 否 | http |
| 104.129.194.44:8800 | ✓ 282ms | ✓ 888ms | ✓ 856ms | 否 | ✓ 983ms | http |
| 198.23.189.151:59394 | ✓ 309ms | 否 | ✓ 896ms | ✓ 1223ms | ✓ 1138ms | http |
| 160.238.65.7:3128 | ✓ 907ms | ✓ 1341ms | ✓ 1283ms | ✓ 1858ms | ✓ 1331ms | http |
| 43.153.182.11:7777 | 否 | 否 | ✓ 1807ms | ✓ 1092ms | ✓ 833ms | http |
| 38.180.2.107:3128 | ✓ 697ms | ✓ 1597ms | ✓ 1643ms | 否 | 否 | http |
| 47.77.216.82:1080 | ✓ 1134ms | ✓ 1278ms | ✓ 470ms | ✓ 1326ms | 否 | http |
| 207.254.71.62:8088 | ✓ 843ms | ✓ 1913ms | ✓ 1439ms | 否 | ✓ 1558ms | http |
| 20.210.76.175:8561 | 否 | 否 | ✓ 701ms | ✓ 1041ms | ✓ 793ms | http |
| 20.18.193.135:8561 | 否 | 否 | ✓ 710ms | ✓ 1089ms | ✓ 763ms | http |
| 20.27.15.49:8561 | 否 | 否 | ✓ 707ms | ✓ 1059ms | ✓ 802ms | http |
| 77.110.119.136:3128 | 否 | 否 | ✓ 305ms | ✓ 1178ms | ✓ 889ms | http |
| 116.171.106.26:3443 | ✓ 1715ms | ✓ 1764ms | ✓ 1682ms | 否 | 否 | http |
| 8.154.21.175:3128 | ✓ 1053ms | ✓ 1437ms | ✓ 981ms | ✓ 1339ms | ✓ 1111ms | http |
| 43.156.132.113:3128 | ✓ 908ms | 否 | ✓ 896ms | ✓ 1280ms | ✓ 1022ms | http |
| 64.188.77.26:3128 | ✓ 1342ms | 否 | ✓ 493ms | 否 | ✓ 1384ms | http |
| 103.155.64.220:8080 | 否 | 否 | ✓ 1258ms | ✓ 1812ms | ✓ 1938ms | http |
| 13.48.13.125:41481 | ✓ 1466ms | 否 | ✓ 1746ms | 否 | ✓ 1865ms | http |
| 195.19.217.200:3128 | ✓ 1575ms | 否 | ✓ 1554ms | 否 | ✓ 1805ms | http |
| 223.16.170.103:80 | ✓ 1354ms | 否 | ✓ 1271ms | ✓ 1407ms | ✓ 1507ms | http |
| 160.238.65.3:3128 | ✓ 567ms | ✓ 1638ms | ✓ 1401ms | 否 | 否 | http |

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
