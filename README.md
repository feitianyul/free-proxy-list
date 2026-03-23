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

最后更新：2026-03-23 00:22:13 UTC（2026-03-23 08:22:13 UTC+8）

**代理总数：86**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 85 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 86 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 43.99.54.236:5555 | ✓ 890ms | ✓ 1180ms | ✓ 822ms | ✓ 1048ms | ✓ 827ms | http |
| 115.231.181.40:8128 | ✓ 853ms | ✓ 1127ms | ✓ 915ms | ✓ 1314ms | ✓ 835ms | http |
| 147.161.210.140:8800 | ✓ 793ms | ✓ 1256ms | ✓ 791ms | ✓ 1500ms | ✓ 1357ms | http |
| 113.160.132.26:8080 | ✓ 1100ms | ✓ 1557ms | ✓ 1056ms | ✓ 1669ms | ✓ 1130ms | http |
| 101.47.73.135:3128 | ✓ 1146ms | 否 | ✓ 1481ms | ✓ 1470ms | ✓ 1387ms | http |
| 167.103.34.108:8800 | ✓ 1166ms | 否 | ✓ 1116ms | ✓ 1348ms | ✓ 1226ms | http |
| 167.103.31.122:8800 | ✓ 1202ms | 否 | ✓ 1187ms | ✓ 1492ms | ✓ 1407ms | http |
| 120.92.212.16:8890 | ✓ 935ms | 否 | ✓ 1586ms | 否 | ✓ 970ms | http |
| 34.101.184.164:3128 | ✓ 1993ms | 否 | ✓ 1421ms | ✓ 1692ms | ✓ 1215ms | http |
| 1.231.81.166:3128 | ✓ 1637ms | ✓ 1211ms | ✓ 983ms | ✓ 1774ms | ✓ 1747ms | http |
| 38.34.179.39:8452 | ✓ 1913ms | ✓ 1505ms | 否 | ✓ 1324ms | 否 | http |
| 217.174.244.117:3129 | ✓ 919ms | ✓ 1717ms | ✓ 1305ms | 否 | 否 | http |
| 138.124.81.12:8888 | ✓ 1183ms | 否 | ✓ 1545ms | ✓ 1698ms | ✓ 1611ms | http |
| 137.220.150.22:6005 | ✓ 1826ms | 否 | 否 | ✓ 1521ms | ✓ 1152ms | http |
| 34.150.20.6:8888 | ✓ 833ms | ✓ 1178ms | ✓ 843ms | ✓ 1100ms | ✓ 833ms | http |
| 101.43.127.100:8877 | ✓ 833ms | ✓ 1059ms | ✓ 972ms | ✓ 1021ms | ✓ 1082ms | http |
| 147.161.239.240:8800 | ✓ 936ms | 否 | ✓ 803ms | ✓ 1437ms | ✓ 1330ms | http |
| 128.199.114.189:9090 | ✓ 1804ms | 否 | ✓ 1796ms | ✓ 1661ms | 否 | http |
| 120.92.212.16:7890 | ✓ 917ms | ✓ 1420ms | 否 | 否 | ✓ 1206ms | http |
| 218.89.134.230:3333 | ✓ 1357ms | ✓ 1530ms | ✓ 1487ms | ✓ 1518ms | ✓ 1121ms | http |
| 45.149.92.147:5001 | ✓ 829ms | 否 | ✓ 1370ms | ✓ 996ms | ✓ 804ms | http |
| 116.80.49.163:3172 | ✓ 1669ms | 否 | ✓ 1645ms | 否 | ✓ 1813ms | http |
| 47.77.193.180:1080 | 否 | ✓ 1033ms | ✓ 907ms | ✓ 917ms | ✓ 701ms | http |
| 106.14.203.63:3333 | ✓ 803ms | 否 | ✓ 1331ms | ✓ 1075ms | ✓ 902ms | http |
| 47.101.149.27:9010 | 否 | ✓ 1812ms | 否 | ✓ 1388ms | ✓ 1912ms | http |
| 210.223.44.230:3128 | 否 | 否 | ✓ 1130ms | ✓ 1233ms | ✓ 886ms | http |
| 8.219.97.248:80 | ✓ 1511ms | 否 | ✓ 1139ms | ✓ 1792ms | 否 | http |
| 38.145.208.181:8445 | ✓ 1655ms | 否 | ✓ 1379ms | ✓ 1300ms | ✓ 1873ms | http |
| 112.163.160.93:3128 | ✓ 1042ms | ✓ 1375ms | ✓ 1043ms | 否 | 否 | http |
| 106.75.15.167:7890 | ✓ 1858ms | ✓ 1398ms | ✓ 933ms | 否 | 否 | http |
| 35.225.22.61:80 | 否 | ✓ 1430ms | ✓ 350ms | 否 | ✓ 1073ms | http |
| 180.125.216.109:8118 | 否 | 否 | ✓ 1039ms | ✓ 1221ms | ✓ 937ms | http |
| 180.103.19.53:1080 | ✓ 1559ms | ✓ 1526ms | ✓ 1580ms | ✓ 1802ms | ✓ 1454ms | http |
| 38.145.208.235:8453 | ✓ 1324ms | ✓ 953ms | ✓ 293ms | ✓ 1199ms | ✓ 856ms | http |
| 185.115.74.185:8080 | ✓ 1100ms | ✓ 1857ms | ✓ 1481ms | 否 | 否 | http |
| 43.165.195.107:3128 | ✓ 1766ms | 否 | ✓ 1587ms | ✓ 1413ms | 否 | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 827ms | ✓ 1014ms | ✓ 1230ms | http |
| 45.93.30.177:6005 | ✓ 1745ms | 否 | 否 | ✓ 1659ms | ✓ 1375ms | http |
| 5.102.109.41:999 | ✓ 1366ms | ✓ 1624ms | ✓ 1906ms | 否 | 否 | http |
| 49.156.44.114:8080 | ✓ 1793ms | ✓ 1878ms | ✓ 1735ms | ✓ 1715ms | 否 | http |
| 121.230.8.11:1080 | ✓ 1251ms | ✓ 1776ms | ✓ 1088ms | ✓ 1734ms | ✓ 1347ms | http |
| 88.80.150.82:8080 | ✓ 1164ms | ✓ 1929ms | ✓ 1979ms | ✓ 1963ms | ✓ 1613ms | https |
| 59.46.216.131:30001 | ✓ 977ms | ✓ 1324ms | ✓ 1002ms | ✓ 1335ms | 否 | http |
| 38.145.208.230:8448 | ✓ 863ms | ✓ 1169ms | ✓ 1602ms | ✓ 1146ms | ✓ 907ms | http |
| 38.34.179.189:8452 | ✓ 843ms | 否 | ✓ 875ms | ✓ 1067ms | 否 | http |
| 38.34.179.14:8444 | ✓ 862ms | ✓ 937ms | ✓ 839ms | 否 | ✓ 1163ms | http |
| 116.236.189.93:29999 | ✓ 937ms | 否 | ✓ 982ms | 否 | ✓ 976ms | http |
| 190.212.131.238:3128 | 否 | 否 | ✓ 895ms | ✓ 1826ms | ✓ 1301ms | http |
| 38.34.183.219:8445 | ✓ 845ms | ✓ 1916ms | 否 | ✓ 928ms | 否 | http |
| 38.34.179.104:8446 | ✓ 939ms | 否 | ✓ 947ms | 否 | ✓ 1256ms | http |
| 45.136.131.39:8451 | ✓ 1491ms | ✓ 1845ms | 否 | ✓ 927ms | 否 | http |
| 38.145.218.76:8445 | ✓ 1572ms | 否 | ✓ 1937ms | ✓ 1867ms | ✓ 1155ms | http |
| 45.136.130.251:8445 | ✓ 1817ms | 否 | 否 | ✓ 1585ms | ✓ 1253ms | http |
| 38.34.179.98:8453 | ✓ 922ms | 否 | ✓ 299ms | ✓ 975ms | ✓ 1999ms | http |
| 38.34.179.86:8452 | ✓ 472ms | ✓ 1110ms | ✓ 1185ms | ✓ 1487ms | ✓ 1134ms | http |
| 38.34.183.130:8452 | ✓ 753ms | ✓ 1065ms | 否 | ✓ 1509ms | ✓ 1590ms | http |
| 194.67.99.223:1080 | ✓ 550ms | ✓ 1903ms | 否 | ✓ 1878ms | ✓ 1435ms | http |
| 166.88.55.83:7890 | ✓ 798ms | ✓ 1300ms | ✓ 803ms | ✓ 1017ms | ✓ 815ms | http |
| 202.141.161.53:30001 | ✓ 1170ms | ✓ 1393ms | ✓ 1205ms | ✓ 1178ms | ✓ 1001ms | http |
| 106.117.208.101:7890 | ✓ 1506ms | 否 | ✓ 1822ms | ✓ 1765ms | ✓ 1224ms | http |
| 45.144.28.81:10808 | ✓ 498ms | 否 | ✓ 1334ms | ✓ 1871ms | ✓ 1161ms | http |
| 38.145.220.33:8448 | ✓ 391ms | ✓ 1573ms | ✓ 897ms | ✓ 942ms | ✓ 1212ms | http |
| 38.34.179.174:8453 | ✓ 666ms | ✓ 1229ms | 否 | ✓ 1204ms | ✓ 1708ms | http |
| 38.34.179.40:8446 | ✓ 1174ms | 否 | ✓ 1491ms | 否 | ✓ 1760ms | http |
| 142.171.224.229:7890 | ✓ 929ms | ✓ 1767ms | ✓ 378ms | ✓ 893ms | ✓ 641ms | http |
| 150.241.77.172:1080 | ✓ 1107ms | 否 | ✓ 1073ms | ✓ 1892ms | ✓ 1468ms | http |
| 45.136.130.192:8451 | ✓ 1080ms | ✓ 1342ms | 否 | ✓ 1489ms | 否 | http |
| 38.145.208.185:8449 | ✓ 1753ms | ✓ 912ms | ✓ 624ms | ✓ 1905ms | ✓ 793ms | http |
| 113.176.92.71:3128 | ✓ 1561ms | ✓ 1569ms | ✓ 1266ms | ✓ 1684ms | ✓ 1177ms | http |
| 103.39.51.190:8080 | ✓ 1955ms | 否 | 否 | ✓ 1894ms | ✓ 1781ms | http |
| 36.37.180.40:8080 | ✓ 1829ms | 否 | ✓ 1969ms | 否 | ✓ 1833ms | http |
| 38.145.218.228:8447 | ✓ 846ms | 否 | ✓ 828ms | ✓ 1282ms | 否 | http |
| 38.34.178.155:8448 | ✓ 1871ms | ✓ 1494ms | 否 | ✓ 1293ms | ✓ 897ms | http |
| 38.145.208.176:8446 | ✓ 1509ms | ✓ 1537ms | 否 | ✓ 1114ms | ✓ 1872ms | http |
| 172.212.68.37:3128 | ✓ 1479ms | ✓ 1379ms | ✓ 1360ms | ✓ 805ms | ✓ 859ms | http |
| 194.87.151.34:1080 | ✓ 1082ms | 否 | ✓ 1528ms | 否 | ✓ 1418ms | http |
| 38.145.208.182:8446 | ✓ 1781ms | ✓ 942ms | ✓ 1171ms | 否 | ✓ 773ms | http |
| 38.145.208.186:8446 | ✓ 1799ms | ✓ 1582ms | 否 | ✓ 1207ms | ✓ 1455ms | http |
| 103.139.138.194:3128 | ✓ 1953ms | 否 | ✓ 1614ms | ✓ 1569ms | ✓ 1504ms | http |
| 38.34.178.154:8445 | ✓ 1820ms | ✓ 1735ms | 否 | ✓ 1371ms | 否 | http |
| 20.120.225.109:3128 | ✓ 967ms | 否 | ✓ 948ms | ✓ 1346ms | ✓ 1026ms | http |
| 103.130.128.13:8080 | ✓ 1953ms | 否 | ✓ 1714ms | ✓ 1851ms | ✓ 1690ms | http |
| 38.34.179.27:8453 | ✓ 886ms | ✓ 1918ms | ✓ 1890ms | ✓ 1147ms | 否 | http |
| 62.113.119.14:8080 | ✓ 1547ms | ✓ 1716ms | ✓ 993ms | ✓ 1908ms | ✓ 1032ms | http |
| 137.220.150.170:6005 | ✓ 1001ms | 否 | ✓ 1585ms | ✓ 1473ms | ✓ 1209ms | http |
| 37.187.109.70:10111 | ✓ 1176ms | 否 | ✓ 1501ms | 否 | ✓ 1781ms | http |

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
