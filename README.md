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

最后更新：2026-04-17 19:59:41 UTC（2026-04-18 03:59:41 UTC+8）

**代理总数：87**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 87 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 87 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 1.231.81.166:3128 | ✓ 1271ms | ✓ 1069ms | ✓ 1402ms | ✓ 960ms | ✓ 853ms | http |
| 113.160.132.26:8080 | ✓ 1512ms | ✓ 1383ms | ✓ 1257ms | ✓ 1294ms | 否 | http |
| 36.141.21.200:7890 | ✓ 1062ms | ✓ 1607ms | ✓ 1435ms | ✓ 1285ms | ✓ 1045ms | http |
| 188.246.224.49:7890 | ✓ 1408ms | ✓ 1814ms | 否 | 否 | ✓ 1875ms | http |
| 35.225.22.61:80 | ✓ 303ms | 否 | ✓ 838ms | 否 | ✓ 994ms | http |
| 149.51.42.10:3128 | ✓ 944ms | ✓ 1906ms | 否 | ✓ 1399ms | 否 | http |
| 34.71.229.255:3128 | ✓ 668ms | ✓ 1360ms | ✓ 947ms | ✓ 1317ms | ✓ 976ms | http |
| 45.149.92.147:5001 | ✓ 1350ms | 否 | ✓ 726ms | ✓ 912ms | ✓ 694ms | http |
| 133.18.123.225:26021 | ✓ 923ms | ✓ 1960ms | 否 | ✓ 1263ms | ✓ 1312ms | http |
| 159.89.191.221:3128 | ✓ 374ms | ✓ 1472ms | ✓ 1566ms | ✓ 1356ms | 否 | http |
| 120.92.108.86:7890 | ✓ 1295ms | 否 | ✓ 1212ms | 否 | ✓ 1574ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1138ms | ✓ 1732ms | 否 | ✓ 1247ms | http |
| 117.236.124.166:3128 | ✓ 1081ms | 否 | ✓ 1500ms | 否 | ✓ 1810ms | http |
| 78.11.96.22:8888 | ✓ 1645ms | ✓ 1984ms | ✓ 1178ms | ✓ 1600ms | ✓ 1693ms | http |
| 34.96.238.40:8080 | ✓ 933ms | 否 | ✓ 1354ms | ✓ 1064ms | 否 | http |
| 116.58.161.203:26021 | ✓ 1407ms | ✓ 1723ms | ✓ 540ms | ✓ 878ms | ✓ 698ms | http |
| 168.144.75.9:3128 | ✓ 1727ms | 否 | ✓ 1107ms | ✓ 1989ms | 否 | http |
| 190.12.150.244:999 | ✓ 1335ms | ✓ 1825ms | ✓ 941ms | ✓ 1731ms | ✓ 1471ms | http |
| 152.32.132.190:7890 | ✓ 853ms | 否 | ✓ 1449ms | 否 | ✓ 760ms | http |
| 120.92.212.16:7890 | ✓ 1087ms | ✓ 1618ms | ✓ 1306ms | ✓ 1417ms | ✓ 1048ms | http |
| 223.84.151.86:30005 | ✓ 1747ms | ✓ 1722ms | ✓ 1526ms | 否 | ✓ 1922ms | http |
| 115.231.181.40:8128 | ✓ 1365ms | 否 | 否 | ✓ 1601ms | ✓ 1913ms | http |
| 218.108.131.186:17890 | 否 | ✓ 1070ms | ✓ 875ms | ✓ 1140ms | ✓ 932ms | http |
| 47.105.98.23:3128 | ✓ 899ms | ✓ 1156ms | ✓ 1016ms | ✓ 1240ms | ✓ 1001ms | http |
| 3.19.213.118:59531 | ✓ 1227ms | 否 | ✓ 842ms | ✓ 1973ms | ✓ 1485ms | http |
| 54.166.0.110:41176 | ✓ 1265ms | 否 | ✓ 808ms | 否 | ✓ 1488ms | http |
| 18.222.132.180:56591 | ✓ 1294ms | 否 | ✓ 921ms | 否 | ✓ 1653ms | http |
| 18.216.7.129:44271 | ✓ 1267ms | 否 | ✓ 1410ms | 否 | ✓ 1900ms | http |
| 18.170.25.193:9002 | ✓ 1239ms | 否 | ✓ 1899ms | 否 | ✓ 1997ms | http |
| 149.51.42.10:8080 | ✓ 949ms | ✓ 1404ms | 否 | ✓ 1395ms | 否 | http |
| 91.233.223.147:3128 | ✓ 876ms | 否 | ✓ 851ms | 否 | ✓ 1621ms | http |
| 103.138.70.165:3129 | 否 | 否 | ✓ 1377ms | ✓ 1555ms | ✓ 1471ms | http |
| 202.141.161.53:10808 | 否 | ✓ 1194ms | ✓ 1780ms | ✓ 1296ms | ✓ 1089ms | http |
| 121.43.196.213:8222 | ✓ 1016ms | ✓ 1095ms | ✓ 871ms | ✓ 1211ms | ✓ 939ms | http |
| 84.47.150.125:1080 | ✓ 834ms | 否 | ✓ 1498ms | 否 | ✓ 1758ms | http |
| 177.93.132.244:3128 | ✓ 949ms | 否 | ✓ 943ms | 否 | ✓ 1769ms | http |
| 8.219.97.248:80 | ✓ 1427ms | 否 | ✓ 1702ms | ✓ 1395ms | 否 | http |
| 43.132.188.134:443 | ✓ 719ms | ✓ 1625ms | 否 | ✓ 860ms | ✓ 1564ms | http |
| 114.237.77.214:1080 | 否 | ✓ 1257ms | ✓ 901ms | ✓ 1302ms | ✓ 977ms | http |
| 52.56.167.111:8906 | ✓ 812ms | 否 | 否 | ✓ 1821ms | ✓ 1419ms | http |
| 34.236.148.220:40687 | ✓ 1927ms | 否 | ✓ 1750ms | ✓ 1962ms | 否 | http |
| 114.237.77.199:1080 | 否 | 否 | ✓ 1013ms | ✓ 1304ms | ✓ 1036ms | http |
| 103.235.67.190:80 | ✓ 1529ms | 否 | ✓ 1184ms | ✓ 1241ms | ✓ 1131ms | http |
| 101.32.244.83:8080 | 否 | 否 | ✓ 956ms | ✓ 1437ms | ✓ 1280ms | http |
| 121.43.196.210:8222 | ✓ 965ms | ✓ 1118ms | ✓ 875ms | ✓ 1124ms | ✓ 957ms | http |
| 114.55.226.123:10086 | ✓ 1182ms | ✓ 1408ms | ✓ 1087ms | ✓ 1369ms | ✓ 1084ms | http |
| 208.87.243.199:7878 | 否 | 否 | ✓ 1000ms | ✓ 975ms | ✓ 992ms | http |
| 13.41.196.179:37858 | ✓ 906ms | 否 | ✓ 1057ms | 否 | ✓ 1564ms | http |
| 52.59.218.12:3538 | ✓ 1459ms | 否 | ✓ 1783ms | 否 | ✓ 1275ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1229ms | ✓ 1195ms | ✓ 1474ms | http |
| 190.60.60.149:8080 | ✓ 1821ms | 否 | ✓ 1774ms | ✓ 1832ms | ✓ 1678ms | http |
| 12.89.176.82:3128 | ✓ 459ms | ✓ 1226ms | ✓ 1265ms | ✓ 1433ms | ✓ 888ms | http |
| 171.244.130.36:3128 | ✓ 1658ms | 否 | 否 | ✓ 1479ms | ✓ 1438ms | http |
| 101.32.243.189:80 | ✓ 1285ms | ✓ 1951ms | ✓ 1557ms | 否 | ✓ 1259ms | http |
| 103.159.96.195:8082 | 否 | 否 | ✓ 1522ms | ✓ 1527ms | ✓ 1510ms | http |
| 210.223.44.230:3128 | ✓ 798ms | ✓ 1254ms | 否 | ✓ 1658ms | ✓ 721ms | http |
| 212.58.132.5:8888 | ✓ 1163ms | 否 | 否 | ✓ 1567ms | ✓ 1317ms | http |
| 174.114.24.95:3128 | ✓ 1529ms | 否 | ✓ 1225ms | ✓ 1776ms | ✓ 1600ms | http |
| 94.158.219.111:3128 | ✓ 1055ms | 否 | ✓ 1207ms | 否 | ✓ 1972ms | http |
| 211.95.152.50:45046 | ✓ 1037ms | ✓ 1155ms | ✓ 1081ms | ✓ 1338ms | 否 | http |
| 59.46.216.131:30001 | ✓ 977ms | ✓ 1324ms | 否 | ✓ 1413ms | ✓ 1076ms | http |
| 51.95.13.205:32982 | ✓ 1770ms | 否 | ✓ 1224ms | 否 | ✓ 1685ms | http |
| 185.138.116.150:8080 | ✓ 899ms | 否 | ✓ 1986ms | 否 | ✓ 1402ms | http |
| 147.45.167.84:3128 | ✓ 1044ms | 否 | ✓ 1775ms | ✓ 1775ms | ✓ 1693ms | http |
| 54.166.0.110:42223 | ✓ 1995ms | 否 | ✓ 946ms | ✓ 1962ms | ✓ 1478ms | http |
| 3.8.4.205:22246 | ✓ 814ms | 否 | 否 | ✓ 1783ms | ✓ 1730ms | http |
| 13.38.59.232:443 | ✓ 1851ms | 否 | ✓ 924ms | 否 | ✓ 1457ms | http |
| 108.130.79.116:51584 | ✓ 1934ms | 否 | ✓ 1299ms | ✓ 1865ms | 否 | http |
| 158.160.215.167:8124 | ✓ 1244ms | 否 | ✓ 1029ms | 否 | ✓ 1969ms | http |
| 217.76.245.80:999 | ✓ 1894ms | ✓ 1516ms | ✓ 1277ms | ✓ 1796ms | ✓ 1518ms | http |
| 52.59.51.29:57486 | ✓ 1508ms | 否 | ✓ 1073ms | ✓ 1727ms | ✓ 1359ms | http |
| 162.240.154.26:3128 | ✓ 1979ms | ✓ 1514ms | ✓ 1022ms | 否 | 否 | http |
| 45.140.147.82:1081 | ✓ 1010ms | 否 | ✓ 1355ms | ✓ 1753ms | ✓ 1889ms | http |
| 139.227.17.70:17890 | ✓ 918ms | ✓ 1130ms | ✓ 1911ms | ✓ 1112ms | ✓ 913ms | http |
| 45.140.147.82:1082 | ✓ 709ms | ✓ 1909ms | 否 | ✓ 1963ms | ✓ 1607ms | http |
| 103.109.173.170:80 | 否 | 否 | ✓ 1681ms | ✓ 1839ms | ✓ 1470ms | http |
| 150.136.239.172:3128 | 否 | 否 | ✓ 1604ms | ✓ 1541ms | ✓ 930ms | http |
| 118.31.1.154:80 | ✓ 1086ms | ✓ 1148ms | ✓ 894ms | ✓ 1144ms | ✓ 976ms | http |
| 180.103.19.207:1080 | ✓ 1062ms | ✓ 1602ms | ✓ 1094ms | ✓ 1389ms | ✓ 1204ms | http |
| 121.230.8.245:1080 | 否 | ✓ 1418ms | ✓ 1305ms | ✓ 1736ms | ✓ 1180ms | http |
| 121.230.8.49:1080 | 否 | ✓ 1360ms | ✓ 1288ms | ✓ 1642ms | ✓ 1358ms | http |
| 201.144.25.226:3128 | ✓ 1752ms | ✓ 1301ms | 否 | 否 | ✓ 1939ms | http |
| 164.92.148.68:3128 | ✓ 1124ms | ✓ 1759ms | ✓ 1735ms | ✓ 1960ms | ✓ 1125ms | http |
| 16.62.229.137:41511 | ✓ 1253ms | 否 | ✓ 1879ms | 否 | ✓ 1734ms | http |
| 42.101.8.101:8888 | ✓ 1213ms | ✓ 1375ms | 否 | 否 | ✓ 1264ms | http |
| 103.39.51.207:8080 | ✓ 1378ms | 否 | 否 | ✓ 1538ms | ✓ 1492ms | http |
| 160.250.5.22:1 | ✓ 1448ms | 否 | ✓ 1391ms | ✓ 1258ms | ✓ 978ms | http |

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
