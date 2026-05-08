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

最后更新：2026-05-08 18:12:18 UTC（2026-05-09 02:12:18 UTC+8）

**代理总数：84**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 84 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 84 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 107.174.64.143:1080 | ✓ 520ms | 否 | ✓ 1072ms | ✓ 1304ms | ✓ 833ms | http |
| 213.111.146.36:18080 | ✓ 1372ms | 否 | ✓ 1156ms | ✓ 1681ms | ✓ 1242ms | http |
| 103.147.152.12:1080 | ✓ 1147ms | ✓ 1551ms | ✓ 1716ms | ✓ 1639ms | ✓ 1687ms | http |
| 115.231.181.40:8128 | ✓ 1167ms | ✓ 1274ms | 否 | ✓ 1367ms | ✓ 1219ms | http |
| 113.160.132.26:8080 | 否 | 否 | ✓ 1403ms | ✓ 1787ms | ✓ 1470ms | http |
| 86.104.72.220:1081 | ✓ 244ms | ✓ 927ms | ✓ 106ms | ✓ 947ms | 否 | http |
| 107.150.41.226:18080 | ✓ 380ms | ✓ 1327ms | ✓ 836ms | ✓ 1308ms | ✓ 1001ms | http |
| 47.77.216.82:1080 | ✓ 1343ms | ✓ 1193ms | ✓ 571ms | ✓ 1044ms | ✓ 793ms | http |
| 212.224.88.212:443 | ✓ 498ms | ✓ 1891ms | ✓ 1362ms | 否 | ✓ 1265ms | http |
| 158.160.215.167:8127 | ✓ 936ms | ✓ 1698ms | ✓ 1427ms | 否 | ✓ 1775ms | http |
| 120.92.212.16:7890 | ✓ 1158ms | 否 | ✓ 1136ms | 否 | ✓ 1098ms | http |
| 185.125.100.115:40000 | 否 | 否 | ✓ 806ms | ✓ 1648ms | ✓ 1270ms | http |
| 185.221.237.57:8443 | ✓ 1106ms | ✓ 1667ms | ✓ 1569ms | 否 | ✓ 1595ms | http |
| 193.160.209.58:1080 | ✓ 1127ms | 否 | ✓ 978ms | 否 | ✓ 1510ms | http |
| 79.137.205.44:40000 | ✓ 1121ms | 否 | ✓ 1358ms | ✓ 1849ms | 否 | http |
| 91.242.229.129:8092 | ✓ 1423ms | ✓ 1847ms | ✓ 834ms | ✓ 1453ms | ✓ 1743ms | http |
| 185.221.237.57:443 | ✓ 478ms | ✓ 1862ms | ✓ 1224ms | ✓ 1745ms | ✓ 1508ms | http |
| 8.219.97.248:80 | ✓ 1043ms | 否 | 否 | ✓ 1840ms | ✓ 1893ms | http |
| 223.84.151.86:30005 | ✓ 1414ms | ✓ 1671ms | ✓ 1718ms | 否 | ✓ 1936ms | http |
| 20.164.75.153:8080 | ✓ 1416ms | 否 | ✓ 996ms | 否 | ✓ 1812ms | http |
| 218.108.131.186:17890 | ✓ 858ms | ✓ 1627ms | 否 | 否 | ✓ 1906ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1655ms | ✓ 1119ms | 否 | ✓ 1321ms | http |
| 45.125.67.37:8443 | 否 | 否 | ✓ 1546ms | ✓ 1330ms | ✓ 1969ms | http |
| 52.56.167.111:55700 | ✓ 987ms | 否 | ✓ 1154ms | ✓ 1738ms | ✓ 1121ms | http |
| 43.156.132.113:3128 | ✓ 905ms | ✓ 1855ms | ✓ 1417ms | ✓ 1255ms | ✓ 977ms | http |
| 51.92.173.133:42274 | ✓ 1039ms | 否 | ✓ 1447ms | 否 | ✓ 1478ms | http |
| 15.161.59.54:443 | ✓ 1148ms | 否 | ✓ 1577ms | 否 | ✓ 1998ms | http |
| 84.47.150.125:1080 | ✓ 1006ms | 否 | 否 | ✓ 1973ms | ✓ 1699ms | http |
| 120.92.212.16:8890 | ✓ 1148ms | 否 | ✓ 1233ms | ✓ 1534ms | ✓ 1097ms | http |
| 206.206.126.177:2412 | ✓ 1525ms | 否 | ✓ 1847ms | 否 | ✓ 1571ms | http |
| 20.210.39.153:8561 | ✓ 714ms | 否 | ✓ 726ms | ✓ 1133ms | ✓ 975ms | http |
| 20.78.118.91:8561 | ✓ 1630ms | 否 | ✓ 804ms | ✓ 990ms | ✓ 817ms | http |
| 20.78.26.206:8561 | ✓ 1769ms | 否 | ✓ 803ms | ✓ 1091ms | ✓ 852ms | http |
| 45.143.94.147:40000 | ✓ 1084ms | 否 | ✓ 1290ms | 否 | ✓ 1627ms | http |
| 116.171.106.26:3443 | ✓ 1568ms | 否 | ✓ 1459ms | ✓ 1814ms | 否 | http |
| 185.195.71.218:18080 | ✓ 1476ms | 否 | 否 | ✓ 1673ms | ✓ 1298ms | http |
| 121.130.199.80:24007 | ✓ 1865ms | ✓ 1845ms | 否 | ✓ 1474ms | ✓ 1556ms | http |
| 38.211.245.18:999 | ✓ 1048ms | 否 | ✓ 1268ms | 否 | ✓ 1941ms | http |
| 62.60.149.161:3128 | ✓ 1665ms | 否 | ✓ 666ms | ✓ 1970ms | 否 | http |
| 168.110.52.228:3128 | ✓ 1510ms | 否 | ✓ 787ms | ✓ 988ms | ✓ 1817ms | http |
| 49.65.127.215:3128 | ✓ 1066ms | ✓ 1370ms | ✓ 1131ms | 否 | ✓ 1102ms | http |
| 64.181.254.251:10443 | 否 | ✓ 1211ms | ✓ 733ms | ✓ 983ms | ✓ 1001ms | http |
| 35.194.4.51:3128 | ✓ 362ms | 否 | ✓ 992ms | ✓ 964ms | ✓ 954ms | http |
| 3.19.213.118:43458 | ✓ 1291ms | 否 | ✓ 1248ms | 否 | ✓ 1487ms | http |
| 34.241.123.181:35589 | ✓ 997ms | 否 | ✓ 1567ms | 否 | ✓ 1433ms | http |
| 63.179.134.206:9780 | ✓ 1583ms | 否 | ✓ 678ms | 否 | ✓ 1611ms | http |
| 3.121.130.230:17289 | ✓ 983ms | 否 | ✓ 1542ms | ✓ 1848ms | ✓ 1598ms | http |
| 15.160.116.45:31892 | ✓ 1760ms | 否 | ✓ 754ms | 否 | ✓ 1591ms | http |
| 51.95.13.205:27107 | ✓ 1953ms | 否 | ✓ 918ms | 否 | ✓ 1476ms | http |
| 77.110.107.80:1080 | ✓ 1090ms | 否 | ✓ 1910ms | 否 | ✓ 1384ms | http |
| 54.229.201.146:38493 | ✓ 1821ms | 否 | 否 | ✓ 1861ms | ✓ 1194ms | http |
| 15.161.131.175:59656 | ✓ 1502ms | 否 | ✓ 1833ms | 否 | ✓ 1466ms | http |
| 173.212.245.136:8888 | ✓ 855ms | 否 | ✓ 1874ms | 否 | ✓ 1964ms | http |
| 174.138.171.164:9000 | ✓ 816ms | 否 | ✓ 896ms | ✓ 1807ms | ✓ 1235ms | http |
| 47.105.98.23:3128 | ✓ 867ms | ✓ 1230ms | ✓ 1219ms | ✓ 1202ms | 否 | http |
| 174.138.171.162:35749 | ✓ 1913ms | ✓ 1707ms | ✓ 1886ms | ✓ 1983ms | ✓ 1737ms | http |
| 103.157.200.126:3128 | ✓ 1855ms | 否 | ✓ 1423ms | 否 | ✓ 1412ms | http |
| 157.0.142.246:10057 | ✓ 979ms | ✓ 1232ms | ✓ 1041ms | ✓ 1353ms | ✓ 1034ms | http |
| 116.171.106.78:3443 | 否 | ✓ 1545ms | ✓ 1510ms | ✓ 1802ms | 否 | http |
| 106.10.55.212:1121 | ✓ 1789ms | ✓ 1265ms | 否 | 否 | ✓ 977ms | http |
| 3.101.133.120:80 | 否 | 否 | ✓ 947ms | ✓ 1473ms | ✓ 1365ms | http |
| 8.209.238.110:47701 | ✓ 1246ms | ✓ 1527ms | ✓ 1055ms | ✓ 1093ms | ✓ 1018ms | http |
| 43.99.54.236:5555 | ✓ 890ms | ✓ 1228ms | ✓ 934ms | ✓ 1117ms | ✓ 820ms | http |
| 77.110.107.80:8080 | ✓ 1234ms | ✓ 1655ms | ✓ 406ms | ✓ 1576ms | ✓ 1168ms | http |
| 138.197.68.35:4857 | ✓ 1226ms | 否 | ✓ 493ms | ✓ 1122ms | ✓ 890ms | http |
| 86.104.74.110:1081 | ✓ 998ms | ✓ 1647ms | ✓ 1098ms | ✓ 1794ms | ✓ 1439ms | http |
| 45.186.6.104:3128 | ✓ 1909ms | ✓ 1914ms | ✓ 1844ms | 否 | 否 | http |
| 3.99.158.157:54722 | ✓ 1178ms | 否 | ✓ 1839ms | 否 | ✓ 1850ms | http |
| 86.104.74.110:1082 | ✓ 1133ms | ✓ 1398ms | ✓ 803ms | ✓ 1887ms | ✓ 1387ms | http |
| 64.188.77.26:3128 | ✓ 1056ms | ✓ 1346ms | ✓ 654ms | ✓ 1862ms | 否 | http |
| 223.16.170.103:3128 | 否 | 否 | ✓ 1290ms | ✓ 1307ms | ✓ 1320ms | http |
| 91.233.223.147:3128 | ✓ 1218ms | 否 | ✓ 1701ms | 否 | ✓ 1539ms | http |
| 158.160.215.167:8124 | ✓ 1667ms | ✓ 1601ms | ✓ 757ms | 否 | 否 | http |
| 80.92.204.47:1082 | ✓ 1122ms | 否 | 否 | ✓ 1440ms | ✓ 1323ms | http |
| 194.59.247.34:10808 | ✓ 1948ms | ✓ 1299ms | ✓ 1128ms | 否 | ✓ 1560ms | http |
| 1.231.81.166:3128 | ✓ 1639ms | ✓ 1219ms | ✓ 1982ms | ✓ 1336ms | ✓ 1057ms | http |
| 210.223.44.230:3128 | 否 | 否 | ✓ 1124ms | ✓ 1597ms | ✓ 1742ms | http |
| 167.71.196.178:80 | ✓ 926ms | 否 | ✓ 935ms | ✓ 1279ms | ✓ 987ms | http |
| 154.90.48.209:9090 | ✓ 1154ms | 否 | 否 | ✓ 1417ms | ✓ 1149ms | http |
| 8.154.21.175:3128 | 否 | 否 | ✓ 1064ms | ✓ 1279ms | ✓ 1064ms | http |
| 121.230.9.231:1080 | ✓ 1186ms | 否 | ✓ 1226ms | ✓ 1727ms | ✓ 1269ms | http |
| 174.138.171.162:35010 | ✓ 868ms | 否 | ✓ 1073ms | 否 | ✓ 1237ms | http |
| 52.59.51.29:2940 | ✓ 981ms | 否 | ✓ 1750ms | 否 | ✓ 1569ms | http |
| 103.39.51.207:8080 | ✓ 1642ms | 否 | 否 | ✓ 1989ms | ✓ 1691ms | http |

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
