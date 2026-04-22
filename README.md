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

最后更新：2026-04-22 17:57:47 UTC（2026-04-23 01:57:47 UTC+8）

**代理总数：64**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 64 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 64 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 35.225.22.61:80 | ✓ 1264ms | ✓ 1423ms | ✓ 1267ms | ✓ 1204ms | ✓ 1046ms | http |
| 1.231.81.166:3128 | ✓ 1703ms | 否 | ✓ 1040ms | ✓ 1113ms | ✓ 874ms | http |
| 152.42.208.139:8118 | ✓ 918ms | 否 | ✓ 1332ms | ✓ 1540ms | ✓ 1012ms | http |
| 46.101.95.183:8888 | ✓ 1394ms | 否 | ✓ 1891ms | 否 | ✓ 1698ms | http |
| 212.58.132.5:8888 | ✓ 1248ms | 否 | ✓ 1448ms | ✓ 1580ms | ✓ 1274ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1630ms | ✓ 1355ms | ✓ 1896ms | ✓ 1183ms | http |
| 152.70.91.193:40000 | ✓ 1746ms | 否 | 否 | ✓ 1848ms | ✓ 1547ms | http |
| 130.61.174.200:1080 | ✓ 540ms | ✓ 1348ms | ✓ 992ms | ✓ 1287ms | ✓ 1160ms | http |
| 207.254.71.62:8088 | ✓ 566ms | 否 | ✓ 1410ms | ✓ 1766ms | ✓ 1506ms | http |
| 218.108.131.186:17890 | ✓ 1033ms | ✓ 1240ms | ✓ 1033ms | ✓ 1326ms | ✓ 1104ms | http |
| 34.96.238.40:8080 | ✓ 1061ms | ✓ 1334ms | ✓ 1707ms | ✓ 1203ms | 否 | http |
| 168.144.75.9:3128 | 否 | 否 | ✓ 1121ms | ✓ 1905ms | ✓ 1093ms | http |
| 45.153.231.229:8080 | ✓ 1844ms | ✓ 1709ms | ✓ 1288ms | 否 | 否 | http |
| 223.84.151.86:30005 | ✓ 1791ms | ✓ 1649ms | ✓ 1344ms | ✓ 1593ms | ✓ 1522ms | http |
| 20.127.128.70:8080 | ✓ 1551ms | 否 | ✓ 1450ms | 否 | ✓ 1849ms | http |
| 161.97.184.191:8080 | ✓ 718ms | ✓ 1858ms | ✓ 1178ms | 否 | ✓ 1764ms | http |
| 137.59.47.73:3128 | ✓ 1774ms | 否 | ✓ 1734ms | ✓ 1332ms | ✓ 1829ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1316ms | ✓ 1110ms | ✓ 1530ms | 否 | http |
| 45.140.147.82:1081 | ✓ 655ms | 否 | 否 | ✓ 1492ms | ✓ 1895ms | http |
| 208.87.243.199:7878 | ✓ 1291ms | 否 | 否 | ✓ 1526ms | ✓ 872ms | http |
| 185.191.236.162:3128 | ✓ 965ms | ✓ 1542ms | ✓ 1949ms | 否 | ✓ 1530ms | http |
| 85.190.99.143:443 | ✓ 1344ms | 否 | ✓ 1640ms | ✓ 1994ms | ✓ 1552ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1635ms | ✓ 1284ms | 否 | ✓ 1284ms | http |
| 34.71.229.255:3128 | ✓ 846ms | ✓ 1364ms | ✓ 1154ms | ✓ 1044ms | ✓ 1241ms | http |
| 120.92.108.86:7890 | ✓ 1401ms | 否 | 否 | ✓ 1964ms | ✓ 1490ms | http |
| 47.101.159.19:8899 | ✓ 1067ms | ✓ 1218ms | ✓ 1112ms | ✓ 1301ms | ✓ 1095ms | http |
| 210.223.44.230:3128 | ✓ 1864ms | ✓ 1157ms | 否 | 否 | ✓ 1422ms | http |
| 47.105.98.23:3128 | ✓ 1045ms | ✓ 1452ms | 否 | ✓ 1435ms | ✓ 1306ms | http |
| 47.95.231.180:8084 | 否 | ✓ 1437ms | ✓ 1047ms | ✓ 1430ms | 否 | http |
| 165.225.72.38:11376 | 否 | ✓ 1822ms | ✓ 972ms | ✓ 1840ms | 否 | http |
| 38.180.2.107:3128 | ✓ 1430ms | ✓ 1935ms | ✓ 1820ms | 否 | 否 | http |
| 121.230.8.227:1080 | ✓ 1390ms | 否 | 否 | ✓ 1481ms | ✓ 1400ms | http |
| 82.114.228.67:1080 | ✓ 1030ms | 否 | ✓ 1942ms | 否 | ✓ 1139ms | http |
| 62.113.119.14:8080 | ✓ 553ms | ✓ 1302ms | ✓ 545ms | ✓ 1495ms | ✓ 1058ms | http |
| 45.140.147.155:1082 | ✓ 1418ms | 否 | ✓ 1961ms | 否 | ✓ 1619ms | http |
| 45.140.147.82:1082 | ✓ 567ms | ✓ 1303ms | ✓ 944ms | ✓ 1708ms | ✓ 1973ms | http |
| 89.208.106.138:10808 | ✓ 864ms | ✓ 1575ms | 否 | 否 | ✓ 1511ms | http |
| 84.47.150.125:1080 | ✓ 1560ms | 否 | ✓ 839ms | 否 | ✓ 1713ms | http |
| 91.99.15.45:2095 | ✓ 1473ms | 否 | ✓ 1497ms | 否 | ✓ 1793ms | http |
| 177.93.132.244:3128 | ✓ 644ms | 否 | ✓ 648ms | 否 | ✓ 1660ms | http |
| 178.156.224.42:3128 | 否 | ✓ 1527ms | ✓ 1925ms | 否 | ✓ 1996ms | http |
| 164.163.42.25:10000 | ✓ 1775ms | 否 | ✓ 1046ms | 否 | ✓ 1948ms | http |
| 168.222.254.136:8888 | ✓ 1347ms | ✓ 1695ms | 否 | 否 | ✓ 1403ms | http |
| 120.92.212.16:7890 | ✓ 1403ms | ✓ 1769ms | ✓ 1814ms | 否 | ✓ 1395ms | http |
| 190.52.110.15:999 | 否 | ✓ 1731ms | ✓ 1215ms | ✓ 1332ms | ✓ 1237ms | http |
| 113.176.92.71:3128 | ✓ 1724ms | ✓ 1572ms | ✓ 1533ms | ✓ 1553ms | ✓ 1246ms | http |
| 121.230.8.162:1080 | ✓ 1378ms | ✓ 1734ms | ✓ 1292ms | ✓ 1875ms | ✓ 1278ms | http |
| 89.111.174.221:8080 | ✓ 765ms | 否 | ✓ 1142ms | 否 | ✓ 1786ms | http |
| 104.129.203.245:10733 | ✓ 496ms | ✓ 1253ms | ✓ 1160ms | ✓ 1190ms | ✓ 1239ms | http |
| 77.232.142.164:3128 | ✓ 757ms | ✓ 1993ms | ✓ 1447ms | ✓ 1964ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1267ms | ✓ 1405ms | ✓ 1412ms | 否 | 否 | http |
| 51.145.178.158:3128 | ✓ 714ms | ✓ 1783ms | ✓ 1862ms | ✓ 1509ms | ✓ 1426ms | http |
| 45.76.207.177:40000 | ✓ 1267ms | 否 | ✓ 1625ms | ✓ 1985ms | ✓ 1921ms | http |
| 91.107.124.215:3128 | ✓ 1083ms | 否 | ✓ 1445ms | 否 | ✓ 1995ms | http |
| 202.141.161.53:10808 | ✓ 1216ms | 否 | ✓ 1301ms | ✓ 1399ms | 否 | http |
| 45.186.6.104:3128 | ✓ 1335ms | ✓ 1858ms | ✓ 1731ms | 否 | 否 | http |
| 198.244.188.127:3128 | ✓ 999ms | ✓ 1883ms | ✓ 1586ms | 否 | ✓ 1805ms | http |
| 42.101.8.101:8888 | ✓ 1261ms | ✓ 1701ms | ✓ 1258ms | ✓ 1556ms | ✓ 1297ms | http |
| 78.11.96.22:8888 | ✓ 841ms | ✓ 1443ms | ✓ 745ms | ✓ 1737ms | ✓ 1285ms | http |
| 51.79.71.106:8080 | ✓ 814ms | ✓ 1977ms | ✓ 1756ms | 否 | ✓ 1196ms | http |
| 8.219.195.129:1080 | ✓ 1386ms | ✓ 1928ms | ✓ 1029ms | ✓ 1295ms | ✓ 1036ms | http |
| 84.47.150.126:1080 | ✓ 1052ms | ✓ 1980ms | ✓ 1671ms | 否 | 否 | http |
| 61.52.131.172:8443 | ✓ 1023ms | ✓ 1412ms | ✓ 1125ms | ✓ 1375ms | ✓ 1108ms | http |
| 121.230.8.136:1080 | ✓ 1523ms | ✓ 1516ms | 否 | 否 | ✓ 1487ms | http |

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
