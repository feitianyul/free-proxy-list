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

最后更新：2026-05-06 21:20:24 UTC（2026-05-07 05:20:24 UTC+8）

**代理总数：85**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 85 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 85 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 47.77.216.82:1080 | ✓ 244ms | ✓ 1173ms | ✓ 954ms | ✓ 929ms | ✓ 596ms | http |
| 8.211.166.184:8081 | ✓ 1208ms | ✓ 1016ms | ✓ 770ms | ✓ 800ms | ✓ 642ms | http |
| 218.108.131.186:17890 | ✓ 787ms | ✓ 1027ms | ✓ 852ms | ✓ 1039ms | ✓ 852ms | http |
| 115.231.181.40:8128 | ✓ 817ms | 否 | ✓ 851ms | ✓ 1079ms | 否 | http |
| 152.32.132.190:7890 | 否 | ✓ 1912ms | 否 | ✓ 799ms | ✓ 640ms | http |
| 113.160.132.26:8080 | ✓ 1376ms | ✓ 1289ms | ✓ 849ms | ✓ 1268ms | ✓ 1118ms | http |
| 59.46.216.131:30001 | ✓ 964ms | 否 | ✓ 1056ms | 否 | ✓ 1108ms | http |
| 91.217.81.131:1080 | ✓ 1805ms | 否 | ✓ 1074ms | 否 | ✓ 1726ms | http |
| 193.160.209.58:1080 | ✓ 1772ms | 否 | ✓ 1110ms | 否 | ✓ 1815ms | http |
| 217.76.245.80:999 | ✓ 1279ms | ✓ 1827ms | ✓ 1488ms | ✓ 1724ms | ✓ 1720ms | http |
| 43.133.44.89:8888 | ✓ 1360ms | 否 | ✓ 692ms | ✓ 1419ms | ✓ 797ms | http |
| 45.125.67.37:8443 | 否 | 否 | ✓ 1047ms | ✓ 1092ms | ✓ 732ms | http |
| 206.206.126.177:2412 | ✓ 1241ms | 否 | ✓ 1970ms | ✓ 1653ms | ✓ 1326ms | http |
| 190.12.150.244:999 | ✓ 1165ms | ✓ 1736ms | ✓ 1066ms | ✓ 1817ms | ✓ 1538ms | http |
| 14.143.222.113:10155 | ✓ 1127ms | 否 | ✓ 969ms | ✓ 1280ms | 否 | http |
| 120.92.108.86:7890 | ✓ 1176ms | 否 | 否 | ✓ 1577ms | ✓ 1409ms | http |
| 20.127.128.70:8080 | ✓ 1772ms | 否 | ✓ 1670ms | 否 | ✓ 1992ms | http |
| 181.119.97.24:999 | ✓ 1110ms | ✓ 1888ms | ✓ 1654ms | ✓ 1897ms | ✓ 1703ms | http |
| 14.143.222.113:57788 | ✓ 1419ms | 否 | ✓ 1710ms | ✓ 1303ms | 否 | http |
| 43.156.132.113:3128 | ✓ 875ms | ✓ 1170ms | ✓ 761ms | ✓ 1004ms | ✓ 783ms | http |
| 8.219.97.248:80 | ✓ 1312ms | 否 | ✓ 1643ms | ✓ 1595ms | 否 | http |
| 120.92.212.16:8890 | ✓ 873ms | ✓ 1105ms | ✓ 968ms | ✓ 1115ms | ✓ 1744ms | http |
| 120.92.212.16:7890 | ✓ 1263ms | ✓ 1118ms | ✓ 1257ms | ✓ 1295ms | ✓ 1184ms | http |
| 45.236.129.64:3128 | ✓ 1555ms | 否 | ✓ 1787ms | 否 | ✓ 1983ms | http |
| 150.107.140.238:3128 | ✓ 1567ms | 否 | ✓ 931ms | ✓ 1111ms | ✓ 1071ms | http |
| 160.238.65.7:3128 | 否 | 否 | ✓ 1809ms | ✓ 1550ms | ✓ 1377ms | http |
| 116.80.50.99:3172 | ✓ 435ms | 否 | ✓ 588ms | ✓ 1755ms | ✓ 1647ms | http |
| 160.238.65.5:3128 | ✓ 1137ms | ✓ 1570ms | ✓ 1130ms | 否 | ✓ 1869ms | http |
| 103.157.200.126:3128 | ✓ 1948ms | 否 | ✓ 1314ms | 否 | ✓ 1600ms | http |
| 5.252.33.13:2025 | ✓ 1538ms | 否 | ✓ 1564ms | 否 | ✓ 1980ms | http |
| 116.80.96.162:3172 | 否 | 否 | ✓ 1474ms | ✓ 1790ms | ✓ 1634ms | http |
| 62.113.119.14:8080 | ✓ 1216ms | ✓ 1554ms | ✓ 897ms | 否 | ✓ 1236ms | http |
| 47.112.25.109:7890 | 否 | 否 | ✓ 900ms | ✓ 1059ms | ✓ 833ms | http |
| 103.157.117.116:8080 | 否 | 否 | ✓ 1541ms | ✓ 1728ms | ✓ 1694ms | http |
| 223.84.151.86:30005 | ✓ 1203ms | 否 | ✓ 1253ms | ✓ 1294ms | ✓ 1246ms | http |
| 84.47.150.125:1080 | ✓ 1163ms | 否 | ✓ 1816ms | 否 | ✓ 1920ms | http |
| 116.80.49.66:3172 | 否 | ✓ 1936ms | 否 | ✓ 1781ms | ✓ 1635ms | http |
| 64.188.77.26:3128 | ✓ 1301ms | 否 | ✓ 981ms | 否 | ✓ 1534ms | http |
| 8.154.21.175:3128 | ✓ 1070ms | ✓ 976ms | ✓ 841ms | ✓ 1117ms | ✓ 912ms | http |
| 202.35.251.72:8080 | ✓ 1703ms | ✓ 997ms | ✓ 1146ms | ✓ 1166ms | ✓ 871ms | http |
| 121.230.9.161:1080 | ✓ 1047ms | ✓ 1445ms | ✓ 998ms | ✓ 1183ms | ✓ 983ms | http |
| 157.0.142.246:10057 | ✓ 972ms | ✓ 1226ms | ✓ 1031ms | ✓ 1270ms | ✓ 1053ms | http |
| 121.230.8.17:1080 | ✓ 1043ms | ✓ 1441ms | ✓ 1028ms | ✓ 1486ms | ✓ 1073ms | http |
| 116.118.48.147:3128 | 否 | 否 | ✓ 1455ms | ✓ 1171ms | ✓ 943ms | http |
| 121.230.9.231:1080 | ✓ 1134ms | ✓ 1263ms | ✓ 1381ms | ✓ 1883ms | ✓ 1126ms | http |
| 103.209.36.58:8080 | ✓ 1714ms | 否 | 否 | ✓ 1981ms | ✓ 1724ms | http |
| 101.32.244.83:8080 | ✓ 982ms | ✓ 1463ms | ✓ 1019ms | ✓ 1383ms | ✓ 1213ms | http |
| 121.43.196.210:8222 | ✓ 932ms | ✓ 993ms | ✓ 792ms | ✓ 1075ms | ✓ 880ms | http |
| 121.43.196.213:8222 | ✓ 911ms | ✓ 1028ms | ✓ 796ms | ✓ 1097ms | ✓ 878ms | http |
| 80.92.204.47:1081 | 否 | ✓ 1577ms | ✓ 1274ms | ✓ 1930ms | ✓ 1281ms | http |
| 116.171.106.26:3443 | ✓ 1457ms | 否 | 否 | ✓ 1814ms | ✓ 1527ms | http |
| 116.80.49.134:3172 | 否 | 否 | ✓ 1541ms | ✓ 1798ms | ✓ 1630ms | http |
| 45.173.12.140:1994 | ✓ 1050ms | ✓ 1973ms | ✓ 1476ms | ✓ 1797ms | ✓ 1528ms | http |
| 116.171.106.111:3443 | ✓ 1348ms | ✓ 1569ms | 否 | ✓ 1571ms | ✓ 1378ms | http |
| 160.238.65.6:3128 | ✓ 639ms | 否 | ✓ 1423ms | ✓ 1481ms | ✓ 1256ms | http |
| 154.90.48.209:9090 | ✓ 1328ms | 否 | 否 | ✓ 1425ms | ✓ 956ms | http |
| 160.238.65.8:3128 | ✓ 1677ms | 否 | ✓ 1148ms | ✓ 1497ms | ✓ 1211ms | http |
| 195.19.217.200:3128 | ✓ 1705ms | 否 | ✓ 1838ms | 否 | ✓ 1815ms | http |
| 223.16.170.103:80 | ✓ 1950ms | ✓ 1593ms | ✓ 1443ms | ✓ 1041ms | ✓ 1068ms | http |
| 212.58.132.5:8888 | ✓ 1165ms | 否 | ✓ 1334ms | ✓ 1551ms | ✓ 1234ms | http |
| 120.92.211.211:7890 | ✓ 1935ms | 否 | ✓ 864ms | 否 | ✓ 1989ms | http |
| 3.101.133.120:80 | ✓ 679ms | ✓ 1269ms | ✓ 1014ms | ✓ 1101ms | ✓ 896ms | http |
| 117.236.124.166:3128 | ✓ 1090ms | 否 | ✓ 1150ms | 否 | ✓ 1620ms | http |
| 110.172.28.217:3128 | ✓ 1528ms | 否 | 否 | ✓ 1612ms | ✓ 1541ms | http |
| 34.96.238.40:8080 | ✓ 1136ms | 否 | ✓ 855ms | ✓ 1069ms | 否 | http |
| 168.110.52.228:3128 | ✓ 1402ms | 否 | ✓ 875ms | ✓ 1069ms | ✓ 1897ms | http |
| 64.188.77.221:3128 | 否 | ✓ 1610ms | ✓ 973ms | 否 | ✓ 1499ms | http |
| 180.125.216.109:8118 | 否 | 否 | ✓ 1015ms | ✓ 1245ms | ✓ 921ms | http |
| 160.238.65.9:3128 | ✓ 1936ms | ✓ 1581ms | ✓ 1164ms | 否 | 否 | http |
| 77.110.119.136:3128 | ✓ 974ms | ✓ 1245ms | ✓ 1085ms | 否 | 否 | http |
| 77.110.107.80:8080 | ✓ 1122ms | ✓ 1780ms | ✓ 742ms | ✓ 1928ms | 否 | http |
| 223.16.170.103:3128 | 否 | 否 | ✓ 975ms | ✓ 1026ms | ✓ 1331ms | http |
| 121.230.9.160:1080 | ✓ 1436ms | ✓ 1792ms | ✓ 1746ms | ✓ 1513ms | ✓ 1178ms | http |
| 94.131.118.39:1081 | ✓ 1952ms | ✓ 1700ms | ✓ 914ms | ✓ 1898ms | ✓ 1707ms | http |
| 121.230.8.253:1080 | ✓ 1022ms | ✓ 1272ms | ✓ 1036ms | ✓ 1483ms | ✓ 1058ms | http |
| 147.45.178.211:14658 | ✓ 1092ms | ✓ 1954ms | ✓ 1734ms | 否 | ✓ 1998ms | http |
| 103.67.46.225:3125 | 否 | 否 | ✓ 1709ms | ✓ 1672ms | ✓ 1526ms | http |
| 61.52.131.172:8443 | ✓ 899ms | ✓ 1139ms | ✓ 892ms | ✓ 1112ms | ✓ 906ms | http |
| 1.20.169.102:8080 | ✓ 1853ms | 否 | ✓ 1374ms | ✓ 1759ms | ✓ 1403ms | http |
| 103.172.70.173:8080 | ✓ 1217ms | ✓ 1969ms | ✓ 1562ms | ✓ 1489ms | ✓ 1289ms | http |
| 172.208.25.199:3128 | ✓ 1045ms | ✓ 1628ms | ✓ 792ms | 否 | 否 | http |
| 16.18.37.186:57094 | ✓ 1406ms | 否 | ✓ 1834ms | 否 | ✓ 1961ms | http |
| 60.249.94.208:3128 | ✓ 668ms | ✓ 912ms | ✓ 675ms | ✓ 945ms | ✓ 711ms | http |
| 116.171.106.78:3443 | ✓ 1446ms | ✓ 1439ms | ✓ 1434ms | ✓ 1743ms | ✓ 1573ms | http |
| 210.223.44.230:3128 | ✓ 1118ms | 否 | ✓ 1483ms | 否 | ✓ 1968ms | http |

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
