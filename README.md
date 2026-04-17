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

最后更新：2026-04-17 18:42:45 UTC（2026-04-18 02:42:45 UTC+8）

**代理总数：78**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 78 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 78 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 208.87.243.199:7878 | ✓ 976ms | ✓ 1007ms | 否 | ✓ 1013ms | ✓ 922ms | http |
| 1.231.81.166:3128 | ✓ 1391ms | ✓ 1032ms | ✓ 1387ms | ✓ 971ms | ✓ 768ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1334ms | ✓ 992ms | ✓ 1489ms | ✓ 1267ms | http |
| 188.246.224.49:7890 | ✓ 1463ms | ✓ 1677ms | ✓ 1786ms | 否 | ✓ 1709ms | http |
| 152.32.132.190:7890 | 否 | 否 | ✓ 1412ms | ✓ 1879ms | ✓ 1741ms | http |
| 159.89.191.221:3128 | ✓ 317ms | ✓ 1365ms | ✓ 1386ms | ✓ 1139ms | ✓ 887ms | http |
| 149.51.42.10:3128 | ✓ 1538ms | ✓ 1471ms | 否 | ✓ 1407ms | 否 | http |
| 34.96.238.40:8080 | ✓ 1305ms | ✓ 1962ms | ✓ 1302ms | 否 | ✓ 1256ms | http |
| 218.108.131.186:17890 | ✓ 1169ms | ✓ 1288ms | ✓ 1057ms | ✓ 1118ms | ✓ 899ms | http |
| 185.138.116.150:8080 | ✓ 732ms | 否 | ✓ 694ms | 否 | ✓ 1268ms | http |
| 45.149.92.147:5001 | ✓ 1048ms | 否 | ✓ 1094ms | ✓ 1518ms | ✓ 868ms | http |
| 168.144.75.9:3128 | ✓ 1017ms | 否 | ✓ 1011ms | ✓ 1919ms | 否 | http |
| 223.84.151.86:30005 | ✓ 1514ms | ✓ 1622ms | ✓ 1031ms | 否 | 否 | http |
| 120.92.108.86:7890 | ✓ 1338ms | 否 | ✓ 1212ms | ✓ 1692ms | ✓ 1425ms | http |
| 157.230.178.216:8088 | ✓ 1006ms | ✓ 1478ms | ✓ 601ms | ✓ 1465ms | ✓ 1304ms | http |
| 45.140.147.155:1082 | ✓ 610ms | ✓ 1867ms | ✓ 1457ms | ✓ 1924ms | ✓ 1405ms | http |
| 114.237.77.214:1080 | ✓ 1177ms | ✓ 1531ms | ✓ 906ms | 否 | 否 | http |
| 168.110.52.228:3128 | ✓ 1193ms | 否 | ✓ 1185ms | ✓ 983ms | ✓ 795ms | http |
| 177.93.132.244:3128 | ✓ 1548ms | 否 | ✓ 766ms | 否 | ✓ 1746ms | http |
| 84.47.150.125:1080 | ✓ 1298ms | 否 | ✓ 1901ms | 否 | ✓ 1761ms | http |
| 190.12.150.244:999 | ✓ 975ms | 否 | ✓ 931ms | ✓ 1701ms | ✓ 1509ms | http |
| 35.225.22.61:80 | ✓ 282ms | 否 | ✓ 798ms | ✓ 1038ms | 否 | http |
| 116.58.161.203:26021 | ✓ 795ms | 否 | 否 | ✓ 1495ms | ✓ 1108ms | http |
| 43.132.188.134:443 | 否 | ✓ 1406ms | 否 | ✓ 1520ms | ✓ 1118ms | http |
| 212.58.132.5:8888 | ✓ 1366ms | 否 | 否 | ✓ 1856ms | ✓ 1443ms | http |
| 120.92.212.16:7890 | ✓ 1287ms | 否 | ✓ 1446ms | ✓ 1500ms | 否 | http |
| 78.11.96.22:8888 | ✓ 1203ms | 否 | ✓ 1214ms | ✓ 1671ms | ✓ 1293ms | http |
| 120.92.212.16:8890 | ✓ 1120ms | ✓ 1234ms | ✓ 1921ms | ✓ 1375ms | ✓ 1251ms | http |
| 47.100.2.5:2020 | ✓ 856ms | ✓ 1290ms | ✓ 860ms | ✓ 1333ms | ✓ 928ms | http |
| 3.145.112.108:55400 | 否 | 否 | ✓ 1284ms | ✓ 1911ms | ✓ 1607ms | http |
| 149.51.42.10:8080 | ✓ 891ms | ✓ 1404ms | 否 | ✓ 1585ms | 否 | http |
| 133.18.123.225:26021 | ✓ 794ms | ✓ 1399ms | ✓ 778ms | ✓ 1024ms | ✓ 1101ms | http |
| 36.141.21.200:7890 | ✓ 1018ms | ✓ 1285ms | ✓ 1403ms | ✓ 1498ms | ✓ 1386ms | http |
| 3.8.4.205:4423 | ✓ 756ms | 否 | 否 | ✓ 1885ms | ✓ 1289ms | http |
| 210.223.44.230:3128 | ✓ 920ms | 否 | ✓ 1635ms | 否 | ✓ 1683ms | http |
| 117.236.124.166:3128 | ✓ 1399ms | 否 | ✓ 1562ms | 否 | ✓ 1635ms | http |
| 72.56.105.251:3128 | 否 | 否 | ✓ 932ms | ✓ 1642ms | ✓ 1692ms | http |
| 103.113.70.189:1081 | ✓ 497ms | 否 | 否 | ✓ 1591ms | ✓ 1984ms | http |
| 120.92.211.211:7890 | 否 | ✓ 1848ms | ✓ 1946ms | 否 | ✓ 1597ms | http |
| 82.114.228.67:1080 | ✓ 809ms | ✓ 1683ms | ✓ 784ms | 否 | ✓ 1628ms | http |
| 103.138.70.165:3129 | ✓ 1369ms | 否 | ✓ 1585ms | 否 | ✓ 1716ms | http |
| 112.65.132.182:3128 | 否 | 否 | ✓ 956ms | ✓ 1122ms | ✓ 919ms | http |
| 103.97.88.47:3128 | ✓ 1597ms | 否 | 否 | ✓ 1849ms | ✓ 1988ms | http |
| 42.200.76.16:3888 | 否 | 否 | ✓ 832ms | ✓ 1003ms | ✓ 802ms | http |
| 59.46.216.131:30001 | ✓ 1008ms | ✓ 1462ms | ✓ 1151ms | 否 | 否 | http |
| 150.249.255.91:3128 | ✓ 1423ms | ✓ 959ms | ✓ 624ms | ✓ 829ms | ✓ 659ms | http |
| 3.19.213.118:59531 | ✓ 1267ms | 否 | ✓ 1941ms | 否 | ✓ 1702ms | http |
| 13.38.217.179:443 | ✓ 1254ms | 否 | 否 | ✓ 1793ms | ✓ 1259ms | http |
| 13.41.196.179:37858 | ✓ 1283ms | 否 | ✓ 1884ms | 否 | ✓ 1512ms | http |
| 34.246.223.187:3128 | ✓ 1946ms | 否 | ✓ 1978ms | 否 | ✓ 1696ms | http |
| 18.170.25.193:9002 | ✓ 1870ms | 否 | ✓ 1857ms | 否 | ✓ 1752ms | http |
| 47.105.98.23:3128 | ✓ 1884ms | ✓ 1888ms | ✓ 1024ms | 否 | 否 | http |
| 47.112.25.109:7890 | ✓ 1519ms | ✓ 1162ms | ✓ 1216ms | ✓ 1150ms | ✓ 897ms | http |
| 202.129.206.239:3128 | ✓ 1223ms | 否 | 否 | ✓ 1614ms | ✓ 1935ms | http |
| 108.131.109.106:3129 | ✓ 736ms | 否 | ✓ 1922ms | 否 | ✓ 1984ms | http |
| 60.249.94.208:3128 | ✓ 1396ms | ✓ 991ms | ✓ 734ms | ✓ 963ms | ✓ 799ms | http |
| 115.231.181.40:8128 | ✓ 868ms | ✓ 1141ms | ✓ 1246ms | 否 | 否 | http |
| 42.101.8.101:8888 | ✓ 1129ms | ✓ 1380ms | ✓ 1049ms | ✓ 1447ms | ✓ 1608ms | http |
| 34.101.184.164:3128 | ✓ 1937ms | 否 | ✓ 1033ms | ✓ 1478ms | ✓ 1109ms | http |
| 202.141.161.53:10808 | 否 | ✓ 1143ms | ✓ 1320ms | ✓ 1288ms | 否 | http |
| 34.71.229.255:3128 | ✓ 505ms | ✓ 1315ms | ✓ 1037ms | ✓ 1027ms | ✓ 790ms | http |
| 45.140.147.82:1081 | ✓ 1730ms | 否 | ✓ 1664ms | ✓ 1804ms | 否 | http |
| 18.201.114.187:36978 | ✓ 800ms | 否 | 否 | ✓ 1838ms | ✓ 1318ms | http |
| 52.56.167.111:8906 | ✓ 1901ms | 否 | ✓ 696ms | ✓ 1713ms | 否 | http |
| 103.113.70.189:1082 | ✓ 1969ms | 否 | ✓ 488ms | ✓ 1705ms | ✓ 1359ms | http |
| 94.131.118.129:1081 | ✓ 1928ms | ✓ 1418ms | ✓ 1300ms | ✓ 1798ms | ✓ 1633ms | http |
| 52.59.51.29:57486 | ✓ 743ms | 否 | 否 | ✓ 1695ms | ✓ 1220ms | http |
| 171.244.130.36:3128 | ✓ 1576ms | 否 | ✓ 1378ms | 否 | ✓ 1358ms | http |
| 15.160.116.45:3129 | ✓ 813ms | 否 | ✓ 1232ms | 否 | ✓ 1337ms | http |
| 116.80.63.67:7777 | ✓ 1647ms | 否 | ✓ 1526ms | 否 | ✓ 1689ms | http |
| 116.80.95.250:3172 | ✓ 1815ms | 否 | ✓ 1542ms | ✓ 1950ms | 否 | http |
| 217.77.102.18:3128 | ✓ 1509ms | 否 | ✓ 1726ms | 否 | ✓ 1974ms | http |
| 12.89.176.82:3128 | ✓ 1009ms | 否 | 否 | ✓ 1292ms | ✓ 937ms | http |
| 103.67.46.225:3125 | ✓ 1712ms | 否 | ✓ 1862ms | ✓ 1697ms | ✓ 1959ms | http |
| 61.52.131.172:8443 | ✓ 1444ms | ✓ 1139ms | ✓ 1315ms | ✓ 1190ms | ✓ 950ms | http |
| 91.233.223.147:3128 | ✓ 914ms | 否 | ✓ 899ms | 否 | ✓ 1628ms | http |
| 158.160.215.167:8124 | ✓ 1213ms | ✓ 1767ms | 否 | 否 | ✓ 1660ms | http |
| 103.122.65.11:8080 | ✓ 1912ms | 否 | ✓ 1795ms | ✓ 1599ms | ✓ 1494ms | http |

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
