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

最后更新：2026-05-29 20:21:38 UTC（2026-05-30 04:21:38 UTC+8）

**代理总数：88**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 88 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 88 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 104.168.153.19:3128 | ✓ 436ms | ✓ 1380ms | ✓ 1318ms | 否 | ✓ 831ms | http |
| 192.99.8.15:8850 | ✓ 395ms | 否 | ✓ 1957ms | ✓ 1469ms | ✓ 1075ms | http |
| 178.253.23.244:8080 | ✓ 1265ms | 否 | ✓ 663ms | ✓ 1644ms | ✓ 1893ms | http |
| 176.111.37.5:39811 | ✓ 1222ms | ✓ 1750ms | ✓ 736ms | 否 | ✓ 1849ms | http |
| 113.160.132.26:8080 | ✓ 1504ms | ✓ 1494ms | ✓ 911ms | ✓ 1229ms | ✓ 1152ms | http |
| 176.111.37.216:39811 | ✓ 1543ms | ✓ 1973ms | ✓ 669ms | 否 | ✓ 1967ms | http |
| 185.200.188.234:10001 | ✓ 1261ms | 否 | ✓ 1611ms | 否 | ✓ 1765ms | http |
| 152.70.91.193:40000 | ✓ 1811ms | 否 | 否 | ✓ 1709ms | ✓ 1692ms | http |
| 209.182.199.206:3128 | ✓ 932ms | ✓ 1732ms | 否 | ✓ 1810ms | ✓ 1015ms | http |
| 174.137.134.182:2999 | ✓ 1369ms | 否 | 否 | ✓ 1335ms | ✓ 832ms | http |
| 161.153.62.49:1080 | 否 | ✓ 834ms | ✓ 238ms | ✓ 1908ms | ✓ 583ms | http |
| 202.28.194.139:31280 | ✓ 1659ms | 否 | 否 | ✓ 1869ms | ✓ 1950ms | http |
| 160.238.65.3:3128 | ✓ 615ms | 否 | ✓ 1351ms | 否 | ✓ 1601ms | http |
| 160.238.65.5:3128 | ✓ 920ms | ✓ 1782ms | ✓ 665ms | 否 | 否 | http |
| 45.89.106.12:8080 | ✓ 1567ms | 否 | ✓ 1221ms | ✓ 1690ms | ✓ 1643ms | http |
| 14.143.222.113:57718 | ✓ 968ms | 否 | ✓ 959ms | ✓ 1395ms | 否 | http |
| 84.47.150.125:1080 | ✓ 1105ms | 否 | ✓ 1458ms | 否 | ✓ 1714ms | http |
| 137.59.47.73:3128 | ✓ 1409ms | ✓ 1259ms | ✓ 1194ms | 否 | ✓ 996ms | http |
| 185.182.65.64:1080 | ✓ 1246ms | ✓ 1953ms | ✓ 1743ms | 否 | 否 | http |
| 185.17.57.178:8080 | ✓ 589ms | ✓ 1611ms | ✓ 1148ms | 否 | ✓ 1915ms | http |
| 185.17.57.179:8080 | ✓ 591ms | ✓ 1554ms | ✓ 1204ms | 否 | ✓ 1909ms | http |
| 38.123.220.147:999 | ✓ 668ms | ✓ 1412ms | ✓ 934ms | ✓ 1470ms | ✓ 1172ms | http |
| 65.109.178.5:8080 | ✓ 1546ms | 否 | ✓ 1713ms | 否 | ✓ 1851ms | http |
| 185.191.239.248:3128 | 否 | 否 | ✓ 998ms | ✓ 1985ms | ✓ 1385ms | http |
| 59.66.26.63:6382 | ✓ 1992ms | 否 | ✓ 1584ms | ✓ 1926ms | 否 | http |
| 43.155.208.105:80 | ✓ 1337ms | ✓ 1563ms | ✓ 1259ms | ✓ 1414ms | ✓ 1064ms | http |
| 104.194.144.249:80 | 否 | ✓ 1900ms | ✓ 1426ms | ✓ 1576ms | ✓ 1326ms | http |
| 43.128.145.26:1080 | ✓ 1887ms | 否 | 否 | ✓ 1704ms | ✓ 709ms | http |
| 160.238.65.9:3128 | ✓ 1051ms | ✓ 1663ms | 否 | ✓ 1947ms | 否 | http |
| 8.154.21.175:3128 | ✓ 863ms | ✓ 1095ms | ✓ 940ms | ✓ 1124ms | ✓ 978ms | http |
| 111.230.27.213:3128 | ✓ 912ms | 否 | ✓ 1997ms | 否 | ✓ 1599ms | http |
| 47.112.25.109:7890 | ✓ 967ms | ✓ 1535ms | ✓ 1934ms | 否 | 否 | http |
| 121.230.8.138:1080 | ✓ 993ms | ✓ 1362ms | ✓ 973ms | ✓ 1375ms | 否 | http |
| 2.26.87.216:1080 | ✓ 1441ms | 否 | ✓ 1468ms | 否 | ✓ 1908ms | http |
| 94.158.244.245:1080 | ✓ 544ms | 否 | ✓ 789ms | ✓ 1225ms | ✓ 1080ms | http |
| 43.133.22.248:9091 | ✓ 857ms | ✓ 1235ms | ✓ 1373ms | ✓ 1459ms | 否 | http |
| 144.31.233.155:5103 | ✓ 583ms | 否 | 否 | ✓ 1573ms | ✓ 1711ms | http |
| 121.230.8.91:1080 | ✓ 1110ms | ✓ 1332ms | ✓ 934ms | ✓ 1441ms | ✓ 1068ms | http |
| 2.27.50.150:8080 | ✓ 1789ms | 否 | ✓ 1517ms | 否 | ✓ 1674ms | http |
| 34.87.80.221:30000 | ✓ 1018ms | ✓ 1837ms | ✓ 891ms | ✓ 1275ms | ✓ 938ms | http |
| 120.132.97.88:7897 | ✓ 1211ms | ✓ 1296ms | ✓ 1333ms | ✓ 1246ms | ✓ 1057ms | http |
| 190.212.131.238:3128 | ✓ 1841ms | 否 | ✓ 1311ms | 否 | ✓ 1967ms | http |
| 57.129.144.178:40000 | ✓ 909ms | 否 | ✓ 1487ms | ✓ 1894ms | ✓ 1775ms | http |
| 170.106.136.181:31002 | ✓ 874ms | ✓ 864ms | ✓ 797ms | ✓ 711ms | ✓ 565ms | http |
| 167.86.95.198:3128 | 否 | ✓ 1972ms | ✓ 1898ms | 否 | ✓ 1467ms | http |
| 121.230.8.22:1080 | 否 | ✓ 1311ms | ✓ 1048ms | ✓ 1326ms | ✓ 1123ms | http |
| 103.157.117.226:81 | ✓ 1860ms | 否 | 否 | ✓ 1383ms | ✓ 1350ms | http |
| 121.230.9.132:1080 | 否 | 否 | ✓ 1785ms | ✓ 1847ms | ✓ 1561ms | http |
| 121.130.177.28:8888 | ✓ 1025ms | 否 | ✓ 1466ms | ✓ 1427ms | ✓ 1338ms | http |
| 101.32.243.189:80 | ✓ 990ms | ✓ 1459ms | 否 | ✓ 1613ms | ✓ 1339ms | http |
| 159.223.41.216:9090 | ✓ 912ms | 否 | ✓ 1156ms | ✓ 1138ms | ✓ 871ms | http |
| 91.132.140.82:3128 | ✓ 695ms | 否 | ✓ 1113ms | 否 | ✓ 1701ms | http |
| 80.150.246.98:443 | ✓ 788ms | 否 | ✓ 1780ms | ✓ 1876ms | ✓ 1165ms | http |
| 128.199.114.189:9090 | ✓ 1118ms | 否 | ✓ 1606ms | ✓ 1972ms | ✓ 1903ms | http |
| 62.113.119.14:8080 | ✓ 729ms | ✓ 1627ms | ✓ 895ms | 否 | 否 | http |
| 193.29.224.20:3128 | ✓ 1394ms | 否 | ✓ 1114ms | ✓ 1920ms | ✓ 1699ms | http |
| 195.25.20.155:3128 | ✓ 1349ms | 否 | ✓ 1943ms | 否 | ✓ 1357ms | http |
| 8.219.97.248:80 | ✓ 1108ms | 否 | ✓ 1144ms | ✓ 1655ms | 否 | http |
| 210.223.44.230:3128 | ✓ 1968ms | 否 | 否 | ✓ 1975ms | ✓ 1961ms | http |
| 2.26.97.98:8080 | ✓ 1094ms | 否 | ✓ 1110ms | 否 | ✓ 1672ms | http |
| 59.66.28.115:6382 | ✓ 1948ms | ✓ 1363ms | ✓ 1354ms | ✓ 1295ms | ✓ 1051ms | http |
| 3.101.133.120:80 | 否 | ✓ 1269ms | ✓ 1094ms | ✓ 1003ms | ✓ 910ms | http |
| 14.143.222.113:10158 | ✓ 1370ms | 否 | ✓ 1267ms | ✓ 1660ms | 否 | http |
| 144.31.73.173:3128 | ✓ 1101ms | 否 | ✓ 1657ms | ✓ 1972ms | 否 | http |
| 46.101.57.56:9050 | ✓ 1110ms | ✓ 1970ms | ✓ 1564ms | 否 | ✓ 1859ms | http |
| 46.101.57.56:9091 | ✓ 1113ms | ✓ 1987ms | ✓ 1545ms | 否 | ✓ 1860ms | http |
| 84.147.118.176:443 | ✓ 1104ms | 否 | 否 | ✓ 1861ms | ✓ 1805ms | http |
| 94.131.118.129:1082 | ✓ 624ms | ✓ 1416ms | ✓ 900ms | ✓ 1517ms | ✓ 1206ms | http |
| 45.129.141.143:3128 | ✓ 1730ms | 否 | ✓ 1812ms | 否 | ✓ 1879ms | http |
| 59.66.16.99:6382 | ✓ 1000ms | ✓ 1262ms | ✓ 1166ms | ✓ 1241ms | ✓ 1021ms | http |
| 103.167.171.149:7778 | 否 | 否 | ✓ 1379ms | ✓ 1446ms | ✓ 1469ms | http |
| 94.131.118.129:1081 | ✓ 691ms | ✓ 1254ms | ✓ 1225ms | 否 | ✓ 1186ms | http |
| 150.107.140.238:3128 | ✓ 1882ms | 否 | ✓ 928ms | ✓ 1463ms | 否 | http |
| 47.84.104.42:3128 | ✓ 778ms | ✓ 1738ms | ✓ 767ms | ✓ 1117ms | ✓ 906ms | http |
| 107.150.97.83:3128 | 否 | ✓ 1140ms | ✓ 958ms | 否 | ✓ 1557ms | http |
| 180.125.216.109:8118 | 否 | ✓ 1123ms | ✓ 931ms | ✓ 1539ms | 否 | http |
| 146.190.80.158:9090 | ✓ 1272ms | 否 | ✓ 818ms | ✓ 1183ms | 否 | http |
| 128.199.121.61:9090 | ✓ 1006ms | 否 | ✓ 1078ms | ✓ 1327ms | ✓ 983ms | http |
| 128.199.116.219:9090 | ✓ 810ms | 否 | 否 | ✓ 1158ms | ✓ 904ms | http |
| 128.199.113.85:9090 | 否 | 否 | ✓ 792ms | ✓ 1176ms | ✓ 931ms | http |
| 152.42.170.187:9090 | ✓ 788ms | 否 | ✓ 788ms | ✓ 1126ms | ✓ 875ms | http |
| 128.199.254.13:9090 | ✓ 997ms | 否 | ✓ 899ms | ✓ 1226ms | ✓ 945ms | http |
| 156.226.173.115:10808 | ✓ 828ms | 否 | 否 | ✓ 1495ms | ✓ 1494ms | http |
| 103.125.174.233:7777 | ✓ 1963ms | 否 | ✓ 1554ms | 否 | ✓ 1493ms | http |
| 61.52.131.172:8443 | 否 | ✓ 1187ms | 否 | ✓ 1732ms | ✓ 1020ms | http |
| 157.245.143.65:7890 | ✓ 823ms | ✓ 1754ms | ✓ 770ms | 否 | 否 | http |
| 103.157.117.116:8080 | 否 | 否 | ✓ 1492ms | ✓ 1779ms | ✓ 1785ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1149ms | ✓ 1743ms | 否 | ✓ 1613ms | http |

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
