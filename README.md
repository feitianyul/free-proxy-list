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

最后更新：2026-04-19 17:39:32 UTC（2026-04-20 01:39:32 UTC+8）

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
| 46.101.95.183:8888 | ✓ 579ms | ✓ 1644ms | ✓ 560ms | ✓ 1499ms | ✓ 1157ms | http |
| 149.51.42.10:3128 | ✓ 902ms | ✓ 1793ms | 否 | ✓ 1942ms | 否 | http |
| 194.104.9.38:3128 | ✓ 530ms | ✓ 1773ms | ✓ 691ms | ✓ 1956ms | ✓ 1417ms | http |
| 91.99.15.45:2095 | ✓ 561ms | ✓ 1292ms | ✓ 1434ms | ✓ 1669ms | ✓ 1649ms | http |
| 81.30.156.115:8080 | ✓ 495ms | ✓ 1470ms | ✓ 1284ms | ✓ 1842ms | ✓ 1460ms | http |
| 188.246.224.49:7890 | ✓ 563ms | 否 | ✓ 1665ms | 否 | ✓ 1707ms | http |
| 152.42.208.139:8118 | ✓ 1638ms | 否 | ✓ 969ms | ✓ 1500ms | ✓ 1010ms | http |
| 185.138.116.150:8080 | 否 | ✓ 1835ms | ✓ 1330ms | ✓ 1858ms | ✓ 1594ms | http |
| 113.160.132.26:8080 | ✓ 1683ms | ✓ 1734ms | 否 | ✓ 1474ms | ✓ 1189ms | http |
| 85.190.99.143:443 | ✓ 753ms | 否 | ✓ 1988ms | 否 | ✓ 1942ms | http |
| 14.247.76.52:8080 | ✓ 1702ms | 否 | 否 | ✓ 1544ms | ✓ 1174ms | http |
| 149.51.42.10:8080 | ✓ 978ms | ✓ 1523ms | 否 | ✓ 1577ms | 否 | http |
| 62.113.119.14:8080 | ✓ 529ms | 否 | ✓ 538ms | ✓ 1396ms | ✓ 1080ms | http |
| 166.88.61.54:8000 | ✓ 1145ms | ✓ 1718ms | ✓ 1158ms | ✓ 1034ms | ✓ 836ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1655ms | ✓ 848ms | ✓ 1118ms | ✓ 934ms | http |
| 177.93.132.244:3128 | ✓ 1507ms | 否 | ✓ 763ms | 否 | ✓ 1618ms | http |
| 59.46.216.131:30001 | ✓ 1277ms | ✓ 1577ms | ✓ 1305ms | 否 | 否 | http |
| 120.92.212.16:7890 | 否 | ✓ 1987ms | 否 | ✓ 1748ms | ✓ 1480ms | http |
| 78.11.96.22:8888 | ✓ 1211ms | ✓ 1508ms | ✓ 1085ms | ✓ 1411ms | ✓ 1582ms | http |
| 35.225.22.61:80 | 否 | ✓ 1299ms | ✓ 1211ms | 否 | ✓ 892ms | http |
| 161.97.184.191:8080 | ✓ 1106ms | 否 | ✓ 1377ms | 否 | ✓ 1566ms | http |
| 120.92.108.86:7890 | ✓ 1776ms | 否 | ✓ 1985ms | 否 | ✓ 1584ms | http |
| 115.231.181.40:8128 | 否 | 否 | ✓ 1923ms | ✓ 1623ms | ✓ 1202ms | http |
| 52.221.190.128:1080 | ✓ 1674ms | ✓ 1313ms | ✓ 1072ms | ✓ 1314ms | ✓ 965ms | http |
| 42.101.8.101:8888 | ✓ 1327ms | ✓ 1614ms | ✓ 1280ms | 否 | 否 | http |
| 20.210.39.153:8561 | ✓ 733ms | ✓ 1481ms | ✓ 974ms | ✓ 1294ms | ✓ 1036ms | http |
| 34.101.184.164:3128 | ✓ 1845ms | 否 | ✓ 1161ms | ✓ 1948ms | ✓ 1172ms | http |
| 45.153.231.229:8080 | ✓ 1137ms | ✓ 1997ms | ✓ 1514ms | 否 | ✓ 1826ms | http |
| 20.78.118.91:8561 | ✓ 1877ms | 否 | ✓ 1447ms | ✓ 1830ms | ✓ 1571ms | http |
| 43.99.54.236:5555 | ✓ 864ms | ✓ 1140ms | ✓ 873ms | ✓ 1090ms | ✓ 812ms | http |
| 218.108.131.186:17890 | ✓ 1018ms | ✓ 1263ms | ✓ 1006ms | ✓ 1292ms | ✓ 1059ms | http |
| 45.76.207.177:40000 | ✓ 1786ms | 否 | ✓ 1024ms | ✓ 1454ms | ✓ 1190ms | http |
| 223.84.151.86:30005 | ✓ 1432ms | ✓ 1457ms | ✓ 1186ms | ✓ 1618ms | ✓ 1421ms | http |
| 117.122.240.82:3338 | ✓ 1365ms | 否 | ✓ 1056ms | ✓ 1342ms | ✓ 1087ms | http |
| 84.47.150.125:1080 | ✓ 922ms | 否 | 否 | ✓ 1802ms | ✓ 1444ms | http |
| 210.223.44.230:3128 | ✓ 751ms | ✓ 1566ms | ✓ 1737ms | ✓ 1095ms | ✓ 923ms | http |
| 47.84.73.61:1080 | ✓ 1037ms | 否 | ✓ 1119ms | ✓ 1263ms | ✓ 1045ms | http |
| 45.140.147.82:1081 | ✓ 1516ms | ✓ 1474ms | ✓ 1195ms | ✓ 1582ms | 否 | http |
| 45.12.151.226:2829 | ✓ 1242ms | 否 | ✓ 1827ms | 否 | ✓ 1650ms | http |
| 94.131.118.129:1081 | ✓ 1240ms | 否 | ✓ 778ms | 否 | ✓ 1685ms | http |
| 195.26.224.49:3128 | 否 | 否 | ✓ 1288ms | ✓ 1493ms | ✓ 1182ms | http |
| 159.89.191.221:3128 | ✓ 1104ms | ✓ 1323ms | 否 | ✓ 974ms | ✓ 889ms | http |
| 152.32.132.190:7890 | ✓ 1554ms | 否 | ✓ 1361ms | 否 | ✓ 868ms | http |
| 208.87.243.199:7878 | 否 | ✓ 1781ms | ✓ 1115ms | ✓ 1717ms | 否 | http |
| 36.141.21.200:7890 | ✓ 1168ms | ✓ 1383ms | ✓ 1163ms | ✓ 1460ms | ✓ 1180ms | http |
| 60.249.94.208:3128 | ✓ 1023ms | ✓ 1432ms | ✓ 890ms | 否 | ✓ 1704ms | http |
| 94.131.118.129:1082 | ✓ 894ms | ✓ 1801ms | ✓ 1482ms | 否 | 否 | http |
| 102.134.49.165:6005 | ✓ 1410ms | ✓ 1466ms | ✓ 1918ms | ✓ 1457ms | ✓ 1060ms | http |
| 164.92.166.127:3128 | ✓ 1807ms | ✓ 1799ms | ✓ 1569ms | ✓ 1567ms | ✓ 1190ms | http |
| 102.134.48.240:6005 | ✓ 1412ms | ✓ 1247ms | ✓ 1647ms | ✓ 1380ms | ✓ 922ms | http |
| 178.213.25.221:7890 | ✓ 1193ms | 否 | ✓ 1161ms | ✓ 1662ms | ✓ 1612ms | http |
| 139.159.97.82:10900 | 否 | 否 | ✓ 1493ms | ✓ 1664ms | ✓ 1386ms | http |
| 120.92.212.16:8890 | ✓ 1119ms | 否 | ✓ 1170ms | ✓ 1458ms | 否 | http |
| 121.230.9.209:1080 | ✓ 1334ms | ✓ 1748ms | ✓ 1351ms | ✓ 1807ms | ✓ 1786ms | http |
| 84.47.150.126:1080 | ✓ 998ms | 否 | ✓ 1335ms | ✓ 1991ms | ✓ 1736ms | http |
| 103.97.140.246:1080 | ✓ 1571ms | 否 | ✓ 1564ms | 否 | ✓ 1817ms | http |
| 103.156.16.12:8818 | ✓ 1570ms | 否 | ✓ 1879ms | ✓ 1640ms | 否 | http |
| 34.96.238.40:8080 | 否 | 否 | ✓ 1235ms | ✓ 1214ms | ✓ 1251ms | http |
| 20.78.26.206:8561 | 否 | 否 | ✓ 1087ms | ✓ 1650ms | ✓ 1128ms | http |
| 218.153.163.186:8890 | 否 | 否 | ✓ 887ms | ✓ 1217ms | ✓ 962ms | http |
| 5.63.111.238:8080 | ✓ 1438ms | 否 | ✓ 1450ms | 否 | ✓ 1884ms | http |
| 207.254.71.62:8088 | ✓ 1621ms | ✓ 1648ms | 否 | 否 | ✓ 1508ms | http |
| 45.140.147.82:1082 | ✓ 987ms | 否 | ✓ 677ms | ✓ 1341ms | ✓ 1136ms | http |
| 47.101.159.19:8899 | ✓ 1064ms | ✓ 1216ms | ✓ 1069ms | 否 | ✓ 1138ms | http |
| 91.193.240.157:9877 | ✓ 1555ms | 否 | ✓ 1253ms | 否 | ✓ 1950ms | http |
| 202.129.206.239:3128 | ✓ 1515ms | 否 | ✓ 1321ms | ✓ 1839ms | 否 | http |
| 116.171.106.15:3443 | ✓ 1611ms | 否 | ✓ 1695ms | 否 | ✓ 1718ms | http |
| 116.171.106.78:3443 | ✓ 1679ms | 否 | 否 | ✓ 1939ms | ✓ 1654ms | http |
| 89.23.123.20:3128 | ✓ 604ms | ✓ 1573ms | ✓ 1094ms | ✓ 1871ms | ✓ 1381ms | http |
| 185.191.236.162:8080 | ✓ 865ms | ✓ 1499ms | ✓ 1694ms | 否 | 否 | http |
| 8.219.97.248:80 | ✓ 1525ms | 否 | ✓ 1533ms | ✓ 1770ms | 否 | http |
| 47.74.226.8:5001 | 否 | ✓ 1746ms | ✓ 1294ms | ✓ 1742ms | ✓ 1510ms | http |
| 192.3.248.190:8014 | 否 | ✓ 1363ms | ✓ 1634ms | ✓ 1675ms | ✓ 1012ms | http |
| 82.148.18.242:443 | ✓ 801ms | 否 | ✓ 1966ms | 否 | ✓ 1732ms | http |
| 45.186.6.104:3128 | ✓ 1695ms | ✓ 1628ms | ✓ 1654ms | 否 | 否 | http |
| 138.124.99.216:8888 | ✓ 787ms | ✓ 1966ms | 否 | ✓ 1836ms | 否 | http |
| 101.32.243.189:80 | ✓ 1433ms | 否 | ✓ 1694ms | ✓ 1804ms | ✓ 1507ms | http |
| 61.52.131.172:8443 | ✓ 1105ms | ✓ 1387ms | ✓ 1110ms | ✓ 1432ms | ✓ 1144ms | http |
| 113.176.92.71:3128 | 否 | ✓ 1678ms | ✓ 1657ms | ✓ 1771ms | 否 | http |
| 138.197.68.35:4857 | 否 | 否 | ✓ 876ms | ✓ 1193ms | ✓ 853ms | http |
| 94.158.219.111:3128 | ✓ 1003ms | ✓ 1928ms | ✓ 879ms | 否 | 否 | http |
| 103.229.126.221:7890 | ✓ 1010ms | ✓ 1716ms | ✓ 1211ms | ✓ 1057ms | ✓ 872ms | http |
| 20.210.76.178:8561 | ✓ 1171ms | ✓ 1016ms | ✓ 749ms | ✓ 1032ms | ✓ 804ms | http |
| 20.210.76.175:8561 | ✓ 1171ms | ✓ 1468ms | ✓ 639ms | ✓ 1024ms | ✓ 838ms | http |
| 103.39.51.207:8080 | ✓ 1653ms | 否 | ✓ 1991ms | ✓ 1736ms | ✓ 1705ms | http |

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
