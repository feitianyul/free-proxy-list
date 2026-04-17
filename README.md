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

最后更新：2026-04-17 17:02:20 UTC（2026-04-18 01:02:20 UTC+8）

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
| 208.87.243.199:7878 | ✓ 967ms | ✓ 1264ms | ✓ 824ms | ✓ 702ms | ✓ 505ms | http |
| 149.51.42.10:3128 | ✓ 761ms | ✓ 1795ms | 否 | ✓ 1619ms | 否 | http |
| 1.231.81.166:3128 | ✓ 1373ms | ✓ 1198ms | ✓ 739ms | ✓ 859ms | ✓ 734ms | http |
| 218.108.131.186:17890 | ✓ 841ms | ✓ 1000ms | ✓ 807ms | ✓ 1085ms | ✓ 891ms | http |
| 36.141.21.200:7890 | ✓ 943ms | ✓ 1772ms | ✓ 995ms | ✓ 1319ms | ✓ 951ms | http |
| 149.51.42.10:8080 | ✓ 760ms | ✓ 1831ms | 否 | ✓ 1590ms | 否 | http |
| 113.160.132.26:8080 | ✓ 1779ms | 否 | ✓ 1170ms | ✓ 1217ms | ✓ 971ms | http |
| 185.138.116.150:8080 | 否 | 否 | ✓ 1558ms | ✓ 1851ms | ✓ 1578ms | http |
| 78.11.96.22:8888 | ✓ 1240ms | ✓ 1818ms | ✓ 1243ms | ✓ 1565ms | ✓ 1383ms | http |
| 34.96.238.40:8080 | 否 | ✓ 1592ms | 否 | ✓ 935ms | ✓ 1358ms | http |
| 177.93.132.244:3128 | 否 | 否 | ✓ 1034ms | ✓ 1849ms | ✓ 1428ms | http |
| 188.246.224.49:7890 | ✓ 723ms | ✓ 1789ms | 否 | ✓ 1969ms | ✓ 1767ms | http |
| 120.92.108.86:7890 | 否 | 否 | ✓ 1317ms | ✓ 1750ms | ✓ 1408ms | http |
| 168.144.75.9:3128 | ✓ 1206ms | 否 | ✓ 1533ms | 否 | ✓ 1551ms | http |
| 190.12.150.244:999 | ✓ 1370ms | 否 | ✓ 1095ms | ✓ 1737ms | ✓ 1446ms | http |
| 43.99.54.236:5555 | ✓ 919ms | ✓ 1085ms | ✓ 790ms | ✓ 808ms | ✓ 651ms | http |
| 117.122.240.82:3338 | ✓ 1185ms | ✓ 1230ms | ✓ 854ms | 否 | 否 | http |
| 202.141.161.53:10808 | ✓ 1018ms | ✓ 1093ms | ✓ 1162ms | ✓ 1216ms | ✓ 1161ms | http |
| 159.89.191.221:3128 | ✓ 287ms | 否 | 否 | ✓ 1436ms | ✓ 1052ms | http |
| 210.223.44.230:3128 | ✓ 1918ms | ✓ 1327ms | 否 | ✓ 1441ms | ✓ 1388ms | http |
| 202.6.200.27:3125 | ✓ 1685ms | 否 | ✓ 1431ms | 否 | ✓ 1706ms | http |
| 84.47.150.125:1080 | ✓ 1101ms | 否 | ✓ 1373ms | 否 | ✓ 1873ms | http |
| 45.140.147.155:1081 | ✓ 1265ms | ✓ 1816ms | ✓ 1886ms | 否 | ✓ 1456ms | http |
| 45.140.147.155:1082 | ✓ 1669ms | ✓ 1745ms | ✓ 1546ms | 否 | ✓ 1458ms | http |
| 212.58.132.5:8888 | ✓ 1725ms | 否 | ✓ 1708ms | ✓ 1459ms | ✓ 1179ms | http |
| 121.230.9.96:1080 | ✓ 1103ms | 否 | ✓ 1176ms | ✓ 1278ms | 否 | http |
| 120.92.212.16:8890 | ✓ 941ms | ✓ 1175ms | 否 | 否 | ✓ 1024ms | http |
| 42.101.8.101:8888 | 否 | 否 | ✓ 972ms | ✓ 1347ms | ✓ 1255ms | http |
| 223.84.151.86:30005 | ✓ 1153ms | ✓ 1159ms | ✓ 960ms | ✓ 1320ms | ✓ 1171ms | http |
| 117.236.124.166:3128 | ✓ 1595ms | 否 | ✓ 1232ms | ✓ 1960ms | 否 | http |
| 120.92.212.16:7890 | ✓ 971ms | ✓ 1946ms | ✓ 1622ms | ✓ 1263ms | 否 | http |
| 157.0.142.246:10057 | ✓ 1553ms | ✓ 1385ms | ✓ 1340ms | ✓ 1790ms | 否 | http |
| 47.121.114.42:3129 | ✓ 1720ms | ✓ 1737ms | ✓ 1801ms | ✓ 1896ms | ✓ 1708ms | http |
| 54.222.174.194:80 | 否 | 否 | ✓ 1746ms | ✓ 1757ms | ✓ 1611ms | http |
| 157.230.178.216:8088 | ✓ 753ms | 否 | ✓ 1177ms | 否 | ✓ 1356ms | http |
| 43.132.188.134:443 | ✓ 1791ms | ✓ 916ms | ✓ 1710ms | 否 | 否 | http |
| 91.233.223.147:3128 | ✓ 941ms | 否 | ✓ 989ms | 否 | ✓ 1626ms | http |
| 116.58.161.203:26021 | ✓ 1520ms | ✓ 1603ms | ✓ 1459ms | ✓ 1136ms | ✓ 1212ms | http |
| 45.140.147.82:1081 | ✓ 906ms | 否 | ✓ 909ms | ✓ 1614ms | ✓ 1278ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 585ms | ✓ 1287ms | ✓ 1062ms | http |
| 34.101.184.164:3128 | ✓ 1339ms | 否 | ✓ 1084ms | ✓ 1379ms | ✓ 1510ms | http |
| 130.61.30.221:8080 | ✓ 894ms | ✓ 1738ms | ✓ 1892ms | ✓ 1805ms | 否 | http |
| 139.159.97.82:10900 | ✓ 1083ms | 否 | ✓ 1526ms | ✓ 1430ms | ✓ 1182ms | http |
| 121.230.8.249:1080 | ✓ 1915ms | ✓ 1607ms | 否 | ✓ 1486ms | ✓ 1359ms | http |
| 121.230.9.209:1080 | ✓ 1187ms | ✓ 1628ms | ✓ 1259ms | ✓ 1595ms | 否 | http |
| 47.84.131.156:8100 | 否 | ✓ 1760ms | ✓ 919ms | 否 | ✓ 874ms | http |
| 47.93.216.160:1081 | ✓ 849ms | ✓ 1156ms | ✓ 912ms | ✓ 1268ms | ✓ 1710ms | http |
| 70.61.188.34:3128 | ✓ 659ms | ✓ 1584ms | ✓ 1572ms | ✓ 1598ms | ✓ 1133ms | http |
| 51.79.71.106:8080 | ✓ 1735ms | 否 | ✓ 1542ms | 否 | ✓ 1560ms | http |
| 101.32.244.83:8080 | ✓ 955ms | ✓ 1891ms | ✓ 934ms | ✓ 1293ms | ✓ 1242ms | http |
| 121.43.196.213:8222 | ✓ 937ms | ✓ 1048ms | ✓ 844ms | ✓ 1166ms | ✓ 921ms | http |
| 121.43.196.210:8222 | ✓ 954ms | ✓ 1216ms | ✓ 834ms | ✓ 1131ms | ✓ 927ms | http |
| 114.55.226.123:10086 | ✓ 1060ms | ✓ 1583ms | ✓ 1037ms | ✓ 1354ms | ✓ 1067ms | http |
| 62.113.119.14:8080 | ✓ 1349ms | ✓ 1582ms | ✓ 854ms | ✓ 1670ms | ✓ 1290ms | http |
| 3.145.87.184:43995 | ✓ 1898ms | 否 | ✓ 1445ms | 否 | ✓ 1920ms | http |
| 103.138.70.165:3129 | 否 | 否 | ✓ 1485ms | ✓ 1850ms | ✓ 1706ms | http |
| 120.132.52.172:8888 | 否 | 否 | ✓ 1239ms | ✓ 1404ms | ✓ 1054ms | http |
| 8.219.97.248:80 | ✓ 957ms | 否 | ✓ 1195ms | 否 | ✓ 1315ms | http |
| 128.199.121.61:9090 | ✓ 1686ms | 否 | ✓ 1229ms | 否 | ✓ 905ms | http |
| 47.100.2.5:2020 | ✓ 817ms | ✓ 1040ms | ✓ 888ms | ✓ 1081ms | ✓ 864ms | http |
| 104.248.151.93:9090 | 否 | 否 | ✓ 934ms | ✓ 1381ms | ✓ 1093ms | http |
| 162.240.154.26:3128 | 否 | ✓ 1186ms | 否 | ✓ 1102ms | ✓ 1178ms | http |
| 45.140.147.82:1082 | ✓ 663ms | 否 | ✓ 733ms | 否 | ✓ 1109ms | http |
| 213.35.113.150:6969 | ✓ 1954ms | 否 | ✓ 1170ms | ✓ 1401ms | 否 | http |
| 217.76.245.80:999 | ✓ 1846ms | ✓ 1537ms | ✓ 1197ms | 否 | 否 | http |
| 42.200.76.16:3888 | ✓ 649ms | 否 | ✓ 656ms | ✓ 819ms | ✓ 686ms | http |
| 152.32.132.190:7890 | ✓ 942ms | 否 | ✓ 1702ms | ✓ 1217ms | ✓ 1523ms | http |
| 175.194.173.105:3128 | 否 | 否 | ✓ 920ms | ✓ 1048ms | ✓ 968ms | http |
| 223.205.77.52:8080 | 否 | 否 | ✓ 1790ms | ✓ 1887ms | ✓ 1517ms | http |
| 20.127.128.70:8080 | ✓ 1737ms | 否 | ✓ 1207ms | 否 | ✓ 1968ms | http |
| 139.227.17.70:17890 | ✓ 1930ms | ✓ 1364ms | ✓ 1078ms | ✓ 1404ms | ✓ 1217ms | http |
| 47.105.98.23:3128 | ✓ 873ms | ✓ 1686ms | ✓ 1675ms | ✓ 1470ms | ✓ 1163ms | http |
| 16.62.229.137:3129 | ✓ 846ms | 否 | 否 | ✓ 1958ms | ✓ 1647ms | http |
| 15.204.233.75:3128 | ✓ 1577ms | ✓ 1991ms | ✓ 1744ms | 否 | ✓ 1293ms | http |
| 157.0.142.246:10061 | 否 | ✓ 1809ms | ✓ 1193ms | ✓ 1600ms | ✓ 1051ms | http |
| 59.46.216.131:30001 | ✓ 1753ms | 否 | ✓ 1357ms | 否 | ✓ 1243ms | http |
| 65.108.203.37:18080 | ✓ 1606ms | 否 | ✓ 1445ms | 否 | ✓ 1855ms | http |
| 135.125.232.193:3128 | ✓ 1168ms | ✓ 1994ms | ✓ 1746ms | 否 | 否 | http |
| 3.71.26.7:17552 | ✓ 1126ms | 否 | ✓ 690ms | ✓ 1716ms | ✓ 1531ms | http |
| 116.80.64.157:7777 | ✓ 1525ms | 否 | 否 | ✓ 1887ms | ✓ 1656ms | http |
| 133.18.123.225:26021 | ✓ 1640ms | 否 | ✓ 571ms | ✓ 1194ms | ✓ 1513ms | http |
| 185.88.101.111:8060 | ✓ 1291ms | ✓ 1716ms | ✓ 1695ms | 否 | ✓ 1496ms | http |
| 18.171.232.214:3129 | ✓ 1308ms | 否 | 否 | ✓ 1912ms | ✓ 1570ms | http |
| 155.212.188.205:8080 | ✓ 1335ms | 否 | ✓ 1574ms | 否 | ✓ 1644ms | http |
| 103.167.171.151:8097 | ✓ 1779ms | 否 | 否 | ✓ 1452ms | ✓ 1465ms | http |
| 220.243.154.24:10878 | ✓ 1040ms | ✓ 1349ms | ✓ 1803ms | ✓ 1239ms | ✓ 1937ms | http |
| 61.52.131.172:8443 | ✓ 833ms | ✓ 1146ms | ✓ 909ms | ✓ 1171ms | ✓ 965ms | http |
| 54.166.0.110:3037 | ✓ 1043ms | 否 | ✓ 1041ms | ✓ 1935ms | ✓ 1947ms | http |
| 104.168.93.120:8080 | ✓ 1046ms | 否 | ✓ 1030ms | ✓ 1667ms | 否 | http |
| 101.32.243.189:80 | ✓ 1233ms | ✓ 1432ms | ✓ 1494ms | ✓ 1392ms | ✓ 1236ms | http |
| 103.39.51.207:8080 | ✓ 1335ms | 否 | 否 | ✓ 1489ms | ✓ 1938ms | http |
| 160.250.5.22:1 | ✓ 1665ms | 否 | ✓ 1346ms | ✓ 1404ms | ✓ 1052ms | http |
| 115.231.181.40:8128 | ✓ 879ms | 否 | ✓ 1905ms | ✓ 1930ms | 否 | http |

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
