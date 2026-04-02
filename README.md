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

最后更新：2026-04-02 11:45:21 UTC（2026-04-02 19:45:21 UTC+8）

**代理总数：76**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 76 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 76 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 43.99.54.236:5555 | ✓ 1503ms | ✓ 1059ms | ✓ 757ms | ✓ 987ms | ✓ 790ms | http |
| 203.80.138.81:50000 | ✓ 1201ms | ✓ 1488ms | ✓ 1181ms | ✓ 1119ms | ✓ 1015ms | http |
| 147.161.210.140:8800 | ✓ 878ms | 否 | ✓ 1100ms | ✓ 1228ms | ✓ 1094ms | http |
| 1.231.81.166:3128 | ✓ 978ms | ✓ 1386ms | ✓ 1749ms | ✓ 1250ms | ✓ 1141ms | http |
| 95.213.217.168:52004 | ✓ 699ms | ✓ 1491ms | 否 | ✓ 1829ms | ✓ 1380ms | http |
| 167.103.115.102:8800 | ✓ 1545ms | 否 | ✓ 1112ms | ✓ 1248ms | ✓ 1189ms | http |
| 113.160.132.26:8080 | ✓ 1545ms | ✓ 1567ms | ✓ 1305ms | ✓ 1726ms | ✓ 1129ms | http |
| 167.103.34.108:8800 | ✓ 1609ms | 否 | ✓ 1694ms | ✓ 1534ms | ✓ 1509ms | http |
| 222.184.48.242:22222 | 否 | 否 | ✓ 1668ms | ✓ 1297ms | ✓ 1015ms | http |
| 222.184.48.251:22222 | ✓ 1077ms | 否 | ✓ 1062ms | 否 | ✓ 1112ms | http |
| 45.167.125.21:999 | ✓ 1316ms | ✓ 1797ms | ✓ 1457ms | ✓ 1730ms | ✓ 1460ms | http |
| 164.90.151.28:3128 | ✓ 448ms | ✓ 1460ms | ✓ 1105ms | ✓ 1043ms | ✓ 691ms | http |
| 119.28.156.42:3128 | ✓ 739ms | ✓ 1119ms | ✓ 926ms | ✓ 1015ms | ✓ 798ms | http |
| 8.219.97.248:80 | ✓ 1660ms | 否 | ✓ 1527ms | 否 | ✓ 1456ms | http |
| 167.103.31.122:8800 | ✓ 1699ms | 否 | ✓ 1446ms | ✓ 1644ms | ✓ 1620ms | http |
| 45.12.151.226:2829 | ✓ 1057ms | ✓ 1752ms | ✓ 1043ms | 否 | ✓ 1154ms | http |
| 180.250.219.58:53281 | ✓ 1958ms | 否 | ✓ 1707ms | ✓ 1815ms | 否 | http |
| 167.103.144.127:8800 | ✓ 1818ms | ✓ 1997ms | ✓ 1338ms | ✓ 1589ms | ✓ 1478ms | http |
| 128.199.121.61:9090 | 否 | 否 | ✓ 1062ms | ✓ 1475ms | ✓ 998ms | http |
| 128.199.113.85:9090 | ✓ 1462ms | 否 | ✓ 843ms | ✓ 1201ms | ✓ 991ms | http |
| 147.161.239.240:8800 | ✓ 1126ms | ✓ 1694ms | ✓ 1544ms | ✓ 1520ms | ✓ 1302ms | http |
| 146.190.80.158:9090 | ✓ 1463ms | 否 | ✓ 912ms | ✓ 1168ms | ✓ 1028ms | http |
| 128.199.116.219:9090 | ✓ 1465ms | 否 | ✓ 864ms | ✓ 1209ms | ✓ 967ms | http |
| 159.223.71.162:8080 | ✓ 1459ms | 否 | ✓ 869ms | ✓ 1244ms | ✓ 942ms | http |
| 159.223.71.162:443 | ✓ 1462ms | 否 | ✓ 854ms | ✓ 1261ms | ✓ 978ms | http |
| 101.43.127.100:8877 | ✓ 1081ms | ✓ 1162ms | ✓ 1655ms | ✓ 1355ms | ✓ 1033ms | http |
| 42.96.16.158:1311 | ✓ 1748ms | 否 | ✓ 1103ms | ✓ 1441ms | ✓ 1172ms | http |
| 160.250.5.22:1 | ✓ 1750ms | 否 | ✓ 1352ms | ✓ 1359ms | ✓ 1102ms | http |
| 5.102.109.41:999 | ✓ 1261ms | ✓ 1382ms | 否 | ✓ 1450ms | 否 | http |
| 177.234.217.88:999 | ✓ 1348ms | ✓ 1958ms | ✓ 1845ms | ✓ 1842ms | ✓ 1523ms | http |
| 218.60.0.214:80 | ✓ 1238ms | ✓ 1485ms | ✓ 1218ms | ✓ 1497ms | ✓ 1216ms | http |
| 38.34.179.27:8451 | 否 | ✓ 981ms | ✓ 1358ms | ✓ 1870ms | ✓ 786ms | http |
| 116.254.118.180:80 | 否 | 否 | ✓ 1047ms | ✓ 1378ms | ✓ 1123ms | http |
| 120.92.212.16:7890 | ✓ 1375ms | 否 | ✓ 1360ms | 否 | ✓ 1146ms | http |
| 120.92.212.16:8890 | ✓ 1128ms | 否 | 否 | ✓ 1686ms | ✓ 1459ms | http |
| 59.46.216.131:30001 | ✓ 1487ms | 否 | ✓ 1170ms | 否 | ✓ 1216ms | http |
| 181.78.44.63:999 | ✓ 1739ms | 否 | ✓ 1210ms | ✓ 1627ms | ✓ 1256ms | http |
| 35.225.22.61:80 | 否 | ✓ 1786ms | 否 | ✓ 958ms | ✓ 783ms | http |
| 116.80.95.238:7777 | ✓ 1608ms | 否 | ✓ 1638ms | ✓ 1957ms | 否 | http |
| 62.113.119.14:8080 | ✓ 1152ms | 否 | ✓ 627ms | ✓ 1550ms | ✓ 1146ms | http |
| 34.96.238.40:8080 | ✓ 1265ms | ✓ 1453ms | ✓ 1261ms | ✓ 1119ms | 否 | http |
| 113.176.92.71:3128 | ✓ 1630ms | ✓ 1450ms | ✓ 956ms | 否 | 否 | http |
| 208.87.243.199:7878 | 否 | ✓ 1584ms | ✓ 1246ms | ✓ 1461ms | ✓ 894ms | http |
| 116.171.106.26:3443 | 否 | ✓ 1773ms | ✓ 1675ms | 否 | ✓ 1787ms | http |
| 150.241.71.15:1080 | ✓ 1205ms | ✓ 1632ms | ✓ 1194ms | ✓ 1591ms | ✓ 1033ms | http |
| 43.153.28.68:3128 | ✓ 428ms | 否 | ✓ 429ms | ✓ 868ms | 否 | http |
| 158.160.215.167:8124 | ✓ 1751ms | ✓ 1843ms | 否 | 否 | ✓ 1746ms | http |
| 150.249.255.91:3128 | 否 | 否 | ✓ 794ms | ✓ 1027ms | ✓ 867ms | http |
| 107.174.208.190:3128 | ✓ 505ms | ✓ 1096ms | ✓ 997ms | ✓ 988ms | ✓ 781ms | http |
| 107.173.111.110:7890 | ✓ 682ms | 否 | ✓ 722ms | 否 | ✓ 751ms | http |
| 121.230.8.97:1080 | ✓ 1140ms | 否 | ✓ 1251ms | 否 | ✓ 1553ms | http |
| 192.71.213.85:9812 | ✓ 1428ms | 否 | ✓ 1925ms | ✓ 1895ms | 否 | http |
| 165.232.146.249:3128 | ✓ 1464ms | 否 | ✓ 1304ms | ✓ 1396ms | 否 | http |
| 38.34.183.234:8450 | 否 | ✓ 1643ms | 否 | ✓ 1434ms | ✓ 1461ms | http |
| 104.248.151.93:9090 | ✓ 867ms | 否 | 否 | ✓ 1275ms | ✓ 1007ms | http |
| 128.199.114.189:9090 | ✓ 1107ms | 否 | ✓ 1460ms | ✓ 1249ms | ✓ 932ms | http |
| 45.136.130.248:8452 | 否 | ✓ 1072ms | 否 | ✓ 1018ms | ✓ 1058ms | http |
| 57.128.188.167:9163 | ✓ 1403ms | 否 | ✓ 1474ms | ✓ 1712ms | ✓ 1510ms | http |
| 5.104.87.17:8051 | ✓ 1968ms | 否 | 否 | ✓ 1979ms | ✓ 1523ms | http |
| 106.117.208.101:7890 | ✓ 1104ms | ✓ 1506ms | ✓ 1183ms | ✓ 1614ms | ✓ 1697ms | http |
| 167.160.184.231:6005 | ✓ 866ms | 否 | ✓ 1113ms | ✓ 1301ms | ✓ 975ms | http |
| 209.97.149.157:80 | ✓ 286ms | ✓ 1218ms | ✓ 1033ms | ✓ 1420ms | ✓ 1003ms | http |
| 128.199.254.13:9090 | ✓ 891ms | 否 | 否 | ✓ 1263ms | ✓ 958ms | http |
| 222.184.48.252:22222 | 否 | ✓ 1322ms | ✓ 1310ms | ✓ 1450ms | 否 | http |
| 223.16.170.103:80 | ✓ 1135ms | 否 | ✓ 1471ms | 否 | ✓ 1283ms | http |
| 85.208.108.43:2094 | ✓ 1319ms | 否 | ✓ 590ms | ✓ 1037ms | ✓ 680ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 895ms | ✓ 1537ms | ✓ 1152ms | http |
| 89.110.124.101:8080 | ✓ 485ms | 否 | ✓ 1447ms | ✓ 1886ms | ✓ 1401ms | http |
| 38.145.208.253:8448 | ✓ 506ms | ✓ 1072ms | ✓ 269ms | ✓ 914ms | ✓ 786ms | http |
| 106.10.55.212:1121 | ✓ 860ms | ✓ 1630ms | ✓ 1059ms | ✓ 1507ms | ✓ 1115ms | http |
| 112.137.170.11:9401 | 否 | 否 | ✓ 1623ms | ✓ 1444ms | ✓ 1154ms | http |
| 72.11.150.178:6005 | 否 | 否 | ✓ 741ms | ✓ 1016ms | ✓ 927ms | http |
| 103.39.51.190:8080 | 否 | 否 | ✓ 1922ms | ✓ 1904ms | ✓ 1764ms | http |
| 38.34.179.60:8450 | 否 | 否 | ✓ 923ms | ✓ 945ms | ✓ 1274ms | http |
| 59.8.203.55:80 | 否 | 否 | ✓ 1374ms | ✓ 1142ms | ✓ 874ms | http |
| 212.58.132.5:8888 | 否 | 否 | ✓ 1465ms | ✓ 1592ms | ✓ 1277ms | http |

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
