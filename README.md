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

最后更新：2026-04-30 05:21:46 UTC（2026-04-30 13:21:46 UTC+8）

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
| 216.180.127.45:1080 | ✓ 610ms | ✓ 1207ms | ✓ 1018ms | ✓ 1404ms | ✓ 1116ms | http |
| 218.108.131.186:17890 | ✓ 802ms | ✓ 1008ms | ✓ 976ms | ✓ 1067ms | ✓ 851ms | http |
| 34.71.229.255:3128 | ✓ 769ms | 否 | ✓ 1097ms | ✓ 1665ms | ✓ 931ms | http |
| 1.231.81.166:3128 | ✓ 1922ms | ✓ 905ms | ✓ 1294ms | ✓ 1281ms | ✓ 978ms | http |
| 137.59.47.73:3128 | ✓ 1503ms | ✓ 1199ms | ✓ 1285ms | ✓ 1116ms | ✓ 900ms | http |
| 113.160.132.26:8080 | ✓ 1528ms | ✓ 1436ms | ✓ 899ms | ✓ 1784ms | ✓ 952ms | http |
| 46.101.95.183:8888 | ✓ 1316ms | 否 | ✓ 867ms | ✓ 1957ms | 否 | http |
| 212.58.132.5:8888 | ✓ 1360ms | 否 | ✓ 1764ms | ✓ 1595ms | ✓ 1603ms | http |
| 45.167.124.71:999 | ✓ 1327ms | ✓ 1976ms | ✓ 1480ms | ✓ 1848ms | 否 | http |
| 154.64.232.35:8080 | ✓ 296ms | ✓ 602ms | ✓ 717ms | ✓ 1234ms | ✓ 709ms | http |
| 34.96.238.40:8080 | ✓ 873ms | ✓ 1451ms | ✓ 1202ms | ✓ 970ms | ✓ 954ms | http |
| 168.110.52.228:3128 | ✓ 579ms | ✓ 1746ms | 否 | ✓ 953ms | ✓ 754ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1474ms | 否 | ✓ 1204ms | ✓ 952ms | http |
| 47.85.51.197:1080 | ✓ 1005ms | 否 | 否 | ✓ 1490ms | ✓ 1045ms | http |
| 132.226.235.199:1080 | ✓ 1432ms | 否 | 否 | ✓ 1304ms | ✓ 879ms | http |
| 103.70.114.149:3128 | 否 | 否 | ✓ 1439ms | ✓ 1420ms | ✓ 1446ms | http |
| 120.132.52.172:8888 | 否 | 否 | ✓ 996ms | ✓ 1208ms | ✓ 1790ms | http |
| 139.159.97.82:10900 | ✓ 1095ms | 否 | ✓ 1277ms | ✓ 1426ms | ✓ 1353ms | http |
| 115.231.181.40:8128 | ✓ 955ms | ✓ 1093ms | 否 | 否 | ✓ 1702ms | http |
| 45.140.147.155:1081 | ✓ 608ms | 否 | ✓ 929ms | 否 | ✓ 1327ms | http |
| 120.92.108.86:7890 | 否 | 否 | ✓ 1898ms | ✓ 1727ms | ✓ 1947ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1552ms | ✓ 885ms | ✓ 1415ms | 否 | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1443ms | ✓ 1824ms | ✓ 1241ms | http |
| 80.92.204.47:1081 | ✓ 1178ms | 否 | ✓ 718ms | ✓ 1886ms | ✓ 1255ms | http |
| 77.110.116.224:3128 | 否 | 否 | ✓ 1492ms | ✓ 1809ms | ✓ 1654ms | http |
| 45.140.147.82:1081 | ✓ 1176ms | ✓ 1509ms | 否 | ✓ 1798ms | ✓ 1226ms | http |
| 34.85.118.216:3128 | ✓ 1332ms | 否 | ✓ 1533ms | ✓ 1294ms | ✓ 1203ms | http |
| 159.223.225.118:8888 | 否 | ✓ 1872ms | ✓ 597ms | ✓ 1815ms | 否 | http |
| 13.38.217.179:39465 | ✓ 1767ms | 否 | ✓ 1530ms | 否 | ✓ 1785ms | http |
| 185.230.190.195:3128 | ✓ 1585ms | ✓ 1991ms | ✓ 1758ms | 否 | 否 | http |
| 206.206.126.177:2412 | ✓ 1393ms | 否 | ✓ 1197ms | ✓ 1028ms | ✓ 848ms | http |
| 20.127.128.70:8080 | ✓ 979ms | 否 | ✓ 1372ms | 否 | ✓ 1103ms | http |
| 129.226.81.110:7890 | ✓ 1177ms | 否 | ✓ 1885ms | ✓ 948ms | ✓ 753ms | http |
| 139.5.189.229:8888 | ✓ 1945ms | 否 | ✓ 1948ms | ✓ 1792ms | ✓ 1783ms | http |
| 103.166.159.115:1080 | ✓ 1373ms | 否 | 否 | ✓ 1641ms | ✓ 1355ms | http |
| 172.236.145.31:7890 | ✓ 1483ms | 否 | ✓ 707ms | ✓ 1193ms | ✓ 1084ms | http |
| 94.72.109.214:8888 | ✓ 1393ms | 否 | ✓ 916ms | ✓ 1781ms | ✓ 1378ms | http |
| 103.157.200.126:3128 | ✓ 1795ms | 否 | ✓ 1440ms | ✓ 1953ms | ✓ 1590ms | http |
| 138.197.68.35:4857 | 否 | 否 | ✓ 335ms | ✓ 1583ms | ✓ 1176ms | http |
| 136.244.67.217:1080 | 否 | ✓ 1647ms | ✓ 1727ms | 否 | ✓ 1568ms | http |
| 8.154.21.175:3128 | ✓ 1859ms | ✓ 968ms | ✓ 863ms | ✓ 1039ms | ✓ 876ms | http |
| 185.21.11.140:1080 | ✓ 1986ms | ✓ 1922ms | 否 | 否 | ✓ 1973ms | http |
| 77.110.119.136:3128 | ✓ 1033ms | 否 | ✓ 450ms | ✓ 1271ms | ✓ 1101ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1368ms | ✓ 1656ms | ✓ 1339ms | 否 | http |
| 47.121.114.42:3129 | ✓ 1476ms | ✓ 1572ms | ✓ 1887ms | ✓ 1746ms | ✓ 1476ms | http |
| 38.180.192.119:3128 | ✓ 183ms | ✓ 672ms | ✓ 1000ms | ✓ 638ms | ✓ 459ms | http |
| 210.223.44.230:3128 | ✓ 748ms | ✓ 788ms | 否 | ✓ 881ms | ✓ 819ms | http |
| 49.156.44.114:8080 | ✓ 1475ms | ✓ 1562ms | ✓ 1867ms | ✓ 1395ms | ✓ 1304ms | http |
| 62.113.119.14:8080 | ✓ 950ms | 否 | ✓ 898ms | ✓ 1818ms | ✓ 1357ms | http |
| 45.140.147.155:1082 | ✓ 1271ms | 否 | ✓ 887ms | 否 | ✓ 1417ms | http |
| 20.27.14.220:8561 | 否 | 否 | ✓ 1353ms | ✓ 1611ms | ✓ 1852ms | http |
| 20.27.15.111:8561 | 否 | 否 | ✓ 1373ms | ✓ 1693ms | ✓ 1863ms | http |
| 157.0.142.246:10057 | 否 | 否 | ✓ 1015ms | ✓ 1213ms | ✓ 1004ms | http |
| 101.132.61.121:8888 | ✓ 1263ms | 否 | ✓ 1323ms | ✓ 1420ms | 否 | http |
| 126.209.18.142:8082 | ✓ 1694ms | 否 | 否 | ✓ 1476ms | ✓ 1499ms | http |
| 8.219.97.248:80 | 否 | 否 | ✓ 1268ms | ✓ 1649ms | ✓ 1502ms | http |
| 107.173.160.222:1080 | ✓ 1260ms | ✓ 836ms | ✓ 617ms | ✓ 852ms | ✓ 731ms | http |
| 38.92.10.152:57579 | ✓ 1435ms | ✓ 948ms | ✓ 289ms | ✓ 1324ms | ✓ 654ms | http |
| 185.234.64.65:1082 | 否 | 否 | ✓ 1817ms | ✓ 1688ms | ✓ 1265ms | http |
| 208.87.243.199:7878 | ✓ 987ms | ✓ 1501ms | ✓ 552ms | 否 | ✓ 1040ms | http |
| 16.62.123.236:8088 | ✓ 889ms | 否 | 否 | ✓ 1840ms | ✓ 1788ms | http |
| 20.27.13.35:8561 | ✓ 1712ms | ✓ 816ms | ✓ 500ms | ✓ 876ms | ✓ 632ms | http |
| 20.27.11.248:8561 | ✓ 1346ms | ✓ 1343ms | ✓ 452ms | ✓ 788ms | ✓ 624ms | http |
| 20.210.39.153:8561 | ✓ 1347ms | ✓ 1369ms | ✓ 672ms | ✓ 1015ms | ✓ 739ms | http |
| 20.78.26.206:8561 | ✓ 1346ms | ✓ 1285ms | ✓ 756ms | ✓ 1026ms | ✓ 742ms | http |
| 20.78.118.91:8561 | ✓ 1333ms | 否 | ✓ 649ms | ✓ 863ms | ✓ 596ms | http |
| 103.3.246.71:3128 | ✓ 1469ms | 否 | ✓ 1210ms | ✓ 1229ms | ✓ 989ms | http |
| 185.234.64.62:1082 | ✓ 944ms | 否 | ✓ 1806ms | 否 | ✓ 1471ms | http |
| 217.182.195.221:30000 | ✓ 1042ms | 否 | ✓ 1438ms | 否 | ✓ 1813ms | http |
| 150.249.255.91:3128 | ✓ 617ms | 否 | ✓ 561ms | ✓ 817ms | 否 | http |
| 147.78.0.81:9443 | ✓ 1797ms | 否 | ✓ 1920ms | 否 | ✓ 1875ms | http |
| 57.128.188.167:9172 | ✓ 1871ms | 否 | ✓ 1939ms | 否 | ✓ 1714ms | http |
| 162.240.154.26:3128 | 否 | ✓ 1524ms | ✓ 1431ms | ✓ 1354ms | 否 | http |
| 103.162.54.26:1111 | ✓ 1004ms | 否 | ✓ 1677ms | ✓ 1731ms | ✓ 1327ms | http |
| 61.52.131.172:8443 | ✓ 848ms | ✓ 1530ms | ✓ 907ms | ✓ 1839ms | ✓ 962ms | http |
| 51.92.173.133:8069 | ✓ 1388ms | 否 | ✓ 1696ms | 否 | ✓ 1810ms | http |
| 121.230.8.136:1080 | ✓ 1107ms | ✓ 1947ms | ✓ 1682ms | ✓ 1408ms | ✓ 1132ms | http |
| 103.39.51.207:8080 | ✓ 1899ms | 否 | ✓ 1789ms | ✓ 1440ms | ✓ 1774ms | http |
| 38.92.10.98:20058 | 否 | 否 | ✓ 929ms | ✓ 715ms | ✓ 549ms | http |

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
