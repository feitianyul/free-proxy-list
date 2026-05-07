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

最后更新：2026-05-07 15:11:55 UTC（2026-05-07 23:11:55 UTC+8）

**代理总数：79**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 79 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 79 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 218.108.131.186:17890 | ✓ 1257ms | ✓ 1081ms | ✓ 1329ms | ✓ 1180ms | ✓ 908ms | http |
| 86.104.74.110:1081 | ✓ 1539ms | 否 | ✓ 708ms | ✓ 1994ms | ✓ 1160ms | http |
| 86.104.74.110:1082 | ✓ 1558ms | ✓ 1514ms | ✓ 1125ms | 否 | ✓ 1348ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1454ms | ✓ 903ms | 否 | ✓ 1251ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1496ms | ✓ 1773ms | ✓ 1834ms | ✓ 1867ms | http |
| 206.206.126.177:2412 | 否 | 否 | ✓ 1531ms | ✓ 1835ms | ✓ 1452ms | http |
| 20.18.193.135:8561 | ✓ 1692ms | 否 | ✓ 1403ms | ✓ 1571ms | ✓ 1072ms | http |
| 20.210.76.178:8561 | ✓ 1566ms | ✓ 946ms | ✓ 469ms | ✓ 959ms | ✓ 859ms | http |
| 20.210.76.175:8561 | ✓ 1566ms | ✓ 1961ms | ✓ 456ms | ✓ 1022ms | ✓ 799ms | http |
| 20.78.118.91:8561 | ✓ 1568ms | ✓ 1827ms | ✓ 1213ms | ✓ 1347ms | ✓ 983ms | http |
| 20.210.39.153:8561 | ✓ 1568ms | ✓ 1976ms | ✓ 1201ms | ✓ 1400ms | ✓ 905ms | http |
| 147.161.247.20:8800 | ✓ 1544ms | 否 | ✓ 1484ms | 否 | ✓ 1671ms | http |
| 59.46.216.131:30001 | ✓ 1942ms | 否 | ✓ 1075ms | ✓ 1302ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1332ms | ✓ 1526ms | ✓ 994ms | 否 | ✓ 1955ms | http |
| 91.242.229.129:8092 | 否 | 否 | ✓ 1281ms | ✓ 1801ms | ✓ 1412ms | http |
| 34.96.238.40:8080 | ✓ 859ms | ✓ 1093ms | ✓ 1267ms | 否 | 否 | http |
| 45.153.231.229:8080 | ✓ 747ms | 否 | ✓ 1351ms | 否 | ✓ 1642ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1229ms | ✓ 1230ms | ✓ 1090ms | ✓ 1779ms | http |
| 43.156.132.113:3128 | ✓ 753ms | ✓ 1553ms | ✓ 797ms | ✓ 1041ms | ✓ 823ms | http |
| 45.125.67.37:8443 | ✓ 1260ms | 否 | ✓ 1371ms | ✓ 1051ms | ✓ 1080ms | http |
| 165.225.113.220:8800 | ✓ 927ms | 否 | ✓ 898ms | 否 | ✓ 886ms | http |
| 8.219.97.248:80 | ✓ 1053ms | 否 | 否 | ✓ 1482ms | ✓ 1435ms | http |
| 20.27.11.248:8561 | ✓ 1431ms | 否 | ✓ 1788ms | ✓ 1809ms | ✓ 1338ms | http |
| 20.78.26.206:8561 | ✓ 1243ms | 否 | ✓ 1779ms | ✓ 1895ms | ✓ 1676ms | http |
| 185.191.236.162:3128 | 否 | 否 | ✓ 775ms | ✓ 1777ms | ✓ 1404ms | http |
| 103.147.152.12:1080 | ✓ 1356ms | ✓ 1660ms | ✓ 1146ms | ✓ 1709ms | ✓ 1337ms | http |
| 195.19.217.200:3128 | ✓ 1744ms | 否 | ✓ 1827ms | 否 | ✓ 1702ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1584ms | ✓ 1811ms | ✓ 1393ms | ✓ 936ms | http |
| 209.126.84.232:8888 | ✓ 1565ms | 否 | ✓ 1619ms | ✓ 1705ms | ✓ 1873ms | http |
| 38.211.245.18:999 | ✓ 1178ms | 否 | ✓ 1138ms | 否 | ✓ 1967ms | http |
| 116.171.106.26:3443 | ✓ 1504ms | 否 | ✓ 1429ms | ✓ 1768ms | 否 | http |
| 116.171.106.78:3443 | ✓ 1982ms | 否 | 否 | ✓ 1468ms | ✓ 1870ms | http |
| 45.59.122.132:80 | 否 | ✓ 1973ms | ✓ 802ms | 否 | ✓ 1393ms | http |
| 47.79.39.142:30000 | 否 | ✓ 1833ms | ✓ 1455ms | ✓ 1024ms | ✓ 722ms | http |
| 43.133.44.89:8888 | ✓ 937ms | 否 | 否 | ✓ 1042ms | ✓ 791ms | http |
| 62.133.60.126:24558 | ✓ 824ms | ✓ 1949ms | ✓ 1249ms | ✓ 1977ms | ✓ 1439ms | http |
| 65.108.203.37:18080 | 否 | 否 | ✓ 1658ms | ✓ 1728ms | ✓ 1883ms | http |
| 84.47.150.125:1080 | 否 | 否 | ✓ 1854ms | ✓ 1859ms | ✓ 1399ms | http |
| 20.164.75.153:8080 | ✓ 1971ms | 否 | ✓ 1725ms | 否 | ✓ 1822ms | http |
| 147.45.178.211:14658 | ✓ 1409ms | 否 | ✓ 1096ms | ✓ 1661ms | ✓ 1472ms | http |
| 120.92.108.86:7890 | ✓ 1279ms | 否 | ✓ 1192ms | 否 | ✓ 1340ms | http |
| 86.104.72.219:1081 | ✓ 364ms | ✓ 1663ms | ✓ 604ms | ✓ 1375ms | ✓ 1291ms | http |
| 62.113.119.14:8080 | ✓ 756ms | ✓ 1691ms | ✓ 735ms | ✓ 1600ms | 否 | http |
| 64.188.77.221:3128 | ✓ 747ms | 否 | ✓ 988ms | ✓ 1926ms | 否 | http |
| 120.132.52.172:8888 | 否 | 否 | ✓ 1046ms | ✓ 1220ms | ✓ 1021ms | http |
| 8.154.21.175:3128 | ✓ 1854ms | ✓ 1025ms | ✓ 882ms | ✓ 1082ms | ✓ 898ms | http |
| 86.104.72.220:1082 | 否 | 否 | ✓ 1172ms | ✓ 1625ms | ✓ 879ms | http |
| 20.27.14.220:8561 | ✓ 1336ms | ✓ 1053ms | ✓ 741ms | ✓ 844ms | ✓ 619ms | http |
| 20.27.15.111:8561 | ✓ 1336ms | ✓ 1114ms | ✓ 681ms | ✓ 845ms | ✓ 619ms | http |
| 104.129.194.46:8800 | ✓ 1394ms | ✓ 1183ms | ✓ 1205ms | 否 | ✓ 1675ms | http |
| 147.161.246.37:8800 | ✓ 797ms | ✓ 1938ms | ✓ 1258ms | ✓ 1585ms | ✓ 1422ms | http |
| 157.0.142.246:10057 | ✓ 1266ms | ✓ 1339ms | ✓ 1148ms | ✓ 1440ms | ✓ 1095ms | http |
| 147.161.246.246:8800 | ✓ 813ms | ✓ 1857ms | ✓ 1317ms | 否 | ✓ 1676ms | http |
| 116.118.48.147:3128 | ✓ 1963ms | 否 | ✓ 1413ms | ✓ 1338ms | ✓ 950ms | http |
| 3.101.133.120:80 | 否 | 否 | ✓ 1325ms | ✓ 1198ms | ✓ 991ms | http |
| 20.27.13.35:8561 | ✓ 520ms | ✓ 1154ms | ✓ 585ms | ✓ 790ms | ✓ 613ms | http |
| 116.80.79.244:3172 | 否 | 否 | ✓ 1460ms | ✓ 1794ms | ✓ 1613ms | http |
| 37.187.109.70:10111 | 否 | ✓ 1701ms | ✓ 617ms | 否 | ✓ 1960ms | http |
| 150.107.140.238:3128 | ✓ 1786ms | 否 | ✓ 966ms | ✓ 1304ms | 否 | http |
| 80.92.204.47:1081 | ✓ 1147ms | 否 | ✓ 1349ms | ✓ 1672ms | ✓ 1265ms | http |
| 1.234.153.14:80 | ✓ 1459ms | ✓ 1520ms | ✓ 872ms | ✓ 879ms | ✓ 672ms | http |
| 47.112.25.109:7890 | ✓ 1088ms | 否 | 否 | ✓ 1346ms | ✓ 1106ms | http |
| 45.236.129.64:3128 | ✓ 1447ms | ✓ 1777ms | ✓ 1903ms | 否 | 否 | http |
| 86.104.72.220:1081 | ✓ 979ms | 否 | ✓ 302ms | ✓ 1142ms | ✓ 1065ms | http |
| 38.194.254.134:999 | ✓ 1564ms | ✓ 1566ms | ✓ 1137ms | 否 | 否 | http |
| 77.110.107.80:1080 | ✓ 1330ms | 否 | ✓ 1491ms | ✓ 1971ms | ✓ 1721ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1147ms | ✓ 1188ms | 否 | ✓ 1975ms | http |
| 121.230.9.160:1080 | 否 | ✓ 1960ms | 否 | ✓ 1603ms | ✓ 1949ms | http |
| 220.121.143.33:3128 | ✓ 721ms | 否 | ✓ 758ms | ✓ 1685ms | 否 | http |
| 91.217.81.131:1080 | ✓ 1355ms | 否 | ✓ 1830ms | 否 | ✓ 1738ms | http |
| 178.63.155.151:8888 | ✓ 1330ms | 否 | ✓ 1655ms | 否 | ✓ 1572ms | http |
| 178.63.155.151:8978 | ✓ 1135ms | 否 | ✓ 889ms | 否 | ✓ 1358ms | http |
| 107.173.42.121:7890 | 否 | ✓ 1704ms | ✓ 783ms | ✓ 1819ms | 否 | http |
| 103.156.15.122:8087 | 否 | 否 | ✓ 1209ms | ✓ 1436ms | ✓ 1424ms | http |
| 103.217.216.65:8181 | 否 | 否 | ✓ 1287ms | ✓ 1428ms | ✓ 1870ms | http |
| 61.52.131.172:8443 | ✓ 871ms | ✓ 1109ms | ✓ 1552ms | ✓ 1214ms | 否 | http |
| 103.157.117.116:8080 | ✓ 1658ms | 否 | ✓ 1919ms | ✓ 1881ms | ✓ 1712ms | http |
| 129.213.162.27:17777 | ✓ 1315ms | 否 | ✓ 1976ms | ✓ 1803ms | ✓ 1224ms | http |
| 144.31.159.23:1080 | ✓ 1584ms | 否 | ✓ 1617ms | 否 | ✓ 1795ms | http |

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
