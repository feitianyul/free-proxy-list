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

最后更新：2026-05-03 22:42:27 UTC（2026-05-04 06:42:27 UTC+8）

**代理总数：69**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 69 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 69 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 47.85.51.197:1080 | ✓ 852ms | ✓ 926ms | ✓ 583ms | ✓ 918ms | 否 | http |
| 218.108.131.186:17890 | ✓ 1250ms | ✓ 1263ms | ✓ 1108ms | ✓ 1950ms | ✓ 1127ms | http |
| 113.160.132.26:8080 | ✓ 1764ms | ✓ 1786ms | 否 | ✓ 1554ms | ✓ 1553ms | http |
| 103.3.246.71:3128 | ✓ 1551ms | 否 | ✓ 1534ms | ✓ 1457ms | ✓ 1383ms | http |
| 45.167.124.71:999 | ✓ 1077ms | ✓ 1722ms | ✓ 1368ms | ✓ 1870ms | ✓ 1442ms | http |
| 154.64.232.35:8080 | ✓ 834ms | ✓ 1057ms | ✓ 754ms | ✓ 994ms | ✓ 1193ms | http |
| 45.59.122.132:80 | ✓ 1025ms | ✓ 1715ms | ✓ 847ms | ✓ 1680ms | ✓ 1512ms | http |
| 168.110.52.228:3128 | ✓ 841ms | ✓ 1302ms | 否 | ✓ 1253ms | ✓ 1055ms | http |
| 46.105.190.40:3128 | 否 | ✓ 1229ms | ✓ 724ms | 否 | ✓ 1296ms | http |
| 194.150.220.163:1082 | ✓ 1041ms | ✓ 1249ms | ✓ 1616ms | 否 | ✓ 1691ms | http |
| 1.231.81.166:3128 | ✓ 913ms | ✓ 1409ms | ✓ 1459ms | ✓ 1761ms | ✓ 1267ms | http |
| 45.153.231.229:8080 | ✓ 1078ms | ✓ 1828ms | ✓ 1954ms | 否 | 否 | http |
| 91.233.223.147:3128 | ✓ 1108ms | 否 | ✓ 1775ms | 否 | ✓ 1887ms | http |
| 34.101.184.164:3128 | ✓ 1719ms | 否 | 否 | ✓ 1602ms | ✓ 1256ms | http |
| 152.32.132.190:7890 | ✓ 873ms | ✓ 1737ms | ✓ 916ms | 否 | ✓ 914ms | http |
| 81.26.190.143:1080 | ✓ 1940ms | ✓ 1857ms | ✓ 1753ms | 否 | 否 | http |
| 38.188.247.12:999 | ✓ 565ms | ✓ 1740ms | ✓ 1343ms | ✓ 1395ms | ✓ 1266ms | http |
| 104.128.138.186:1080 | ✓ 828ms | ✓ 1661ms | ✓ 1265ms | ✓ 1815ms | ✓ 1466ms | http |
| 20.127.128.70:8080 | ✓ 1888ms | 否 | ✓ 1746ms | ✓ 1583ms | ✓ 892ms | http |
| 103.157.200.126:3128 | ✓ 1419ms | 否 | ✓ 906ms | 否 | ✓ 1309ms | http |
| 212.58.132.5:8888 | ✓ 1394ms | 否 | ✓ 1634ms | ✓ 1607ms | ✓ 1257ms | http |
| 120.92.108.86:7890 | ✓ 1481ms | 否 | ✓ 1968ms | 否 | ✓ 1766ms | http |
| 162.255.110.107:8080 | ✓ 629ms | 否 | ✓ 1183ms | ✓ 1553ms | 否 | http |
| 46.105.190.38:3128 | ✓ 949ms | ✓ 1675ms | ✓ 923ms | 否 | 否 | http |
| 38.7.23.138:999 | 否 | 否 | ✓ 1230ms | ✓ 1482ms | ✓ 1275ms | http |
| 190.52.110.71:999 | 否 | ✓ 1630ms | ✓ 984ms | ✓ 1654ms | 否 | http |
| 94.131.118.39:1081 | ✓ 927ms | ✓ 1283ms | ✓ 991ms | 否 | ✓ 1224ms | http |
| 20.164.75.153:8080 | ✓ 1162ms | 否 | ✓ 1869ms | 否 | ✓ 1994ms | http |
| 206.206.126.177:2412 | ✓ 883ms | ✓ 1728ms | ✓ 1400ms | ✓ 1231ms | ✓ 951ms | http |
| 125.76.214.178:8091 | ✓ 1071ms | ✓ 1405ms | ✓ 1176ms | 否 | 否 | http |
| 80.92.204.47:1081 | ✓ 902ms | ✓ 1335ms | ✓ 673ms | ✓ 1824ms | ✓ 1069ms | http |
| 185.21.11.140:1080 | ✓ 1405ms | ✓ 1492ms | 否 | ✓ 1491ms | ✓ 1302ms | http |
| 193.123.250.39:1080 | 否 | 否 | ✓ 1780ms | ✓ 1135ms | ✓ 1048ms | http |
| 120.92.212.16:8890 | ✓ 1493ms | 否 | ✓ 1940ms | 否 | ✓ 1989ms | http |
| 223.84.151.86:30005 | ✓ 1630ms | ✓ 1559ms | ✓ 1578ms | ✓ 1999ms | ✓ 1492ms | http |
| 8.211.166.184:8081 | ✓ 663ms | ✓ 970ms | ✓ 838ms | ✓ 1067ms | ✓ 816ms | http |
| 107.173.42.121:7890 | 否 | ✓ 1200ms | ✓ 472ms | ✓ 983ms | 否 | http |
| 8.154.21.175:3128 | ✓ 1041ms | ✓ 1271ms | ✓ 985ms | ✓ 1343ms | ✓ 1106ms | http |
| 45.125.67.37:8443 | ✓ 1203ms | 否 | ✓ 1315ms | ✓ 1361ms | ✓ 1280ms | http |
| 105.159.148.25:5205 | ✓ 884ms | ✓ 1662ms | ✓ 1224ms | ✓ 1552ms | ✓ 1283ms | http |
| 119.195.17.15:3192 | 否 | ✓ 1556ms | ✓ 968ms | ✓ 1628ms | ✓ 1948ms | http |
| 105.159.148.99:5431 | ✓ 939ms | ✓ 1636ms | ✓ 1194ms | ✓ 1578ms | ✓ 1515ms | http |
| 120.92.212.16:7890 | ✓ 1936ms | ✓ 1630ms | ✓ 1358ms | ✓ 1462ms | ✓ 1707ms | http |
| 23.138.80.49:999 | 否 | 否 | ✓ 1430ms | ✓ 1406ms | ✓ 1390ms | http |
| 103.156.248.53:8080 | ✓ 1517ms | 否 | ✓ 1878ms | ✓ 1624ms | ✓ 1666ms | http |
| 110.235.136.71:8081 | 否 | 否 | ✓ 1435ms | ✓ 1661ms | ✓ 1774ms | http |
| 154.56.114.10:8082 | 否 | 否 | ✓ 1874ms | ✓ 1632ms | ✓ 1602ms | http |
| 110.232.85.86:6060 | ✓ 1636ms | 否 | 否 | ✓ 1779ms | ✓ 1741ms | http |
| 103.93.93.221:8181 | 否 | 否 | ✓ 1764ms | ✓ 1757ms | ✓ 1691ms | http |
| 148.230.4.241:999 | ✓ 1106ms | ✓ 1629ms | ✓ 844ms | 否 | 否 | http |
| 154.90.48.209:9090 | ✓ 1730ms | 否 | ✓ 1069ms | 否 | ✓ 1279ms | http |
| 116.171.106.111:3443 | ✓ 1816ms | 否 | ✓ 1672ms | ✓ 1937ms | 否 | http |
| 3.101.133.120:80 | ✓ 533ms | ✓ 1436ms | ✓ 1341ms | ✓ 1303ms | ✓ 1368ms | http |
| 45.140.147.155:1082 | ✓ 1008ms | ✓ 1727ms | ✓ 744ms | ✓ 1750ms | 否 | http |
| 150.136.153.231:80 | ✓ 1977ms | ✓ 1341ms | 否 | 否 | ✓ 1237ms | http |
| 147.45.178.211:14658 | ✓ 977ms | ✓ 1581ms | 否 | ✓ 1765ms | 否 | http |
| 119.195.17.15:3176 | ✓ 936ms | ✓ 1136ms | ✓ 1362ms | ✓ 1843ms | 否 | http |
| 117.236.124.166:3128 | ✓ 1425ms | 否 | ✓ 1072ms | 否 | ✓ 1624ms | http |
| 178.156.224.42:3128 | ✓ 891ms | ✓ 1817ms | 否 | 否 | ✓ 1600ms | http |
| 43.133.44.89:8888 | ✓ 1087ms | 否 | 否 | ✓ 1527ms | ✓ 1678ms | http |
| 38.252.213.69:999 | ✓ 1954ms | 否 | 否 | ✓ 1860ms | ✓ 1935ms | http |
| 45.174.77.1:999 | ✓ 702ms | ✓ 1368ms | ✓ 981ms | 否 | 否 | http |
| 59.46.216.131:30001 | ✓ 1227ms | ✓ 1512ms | ✓ 1330ms | ✓ 1506ms | 否 | http |
| 152.42.177.32:8888 | ✓ 1275ms | 否 | ✓ 1280ms | ✓ 1556ms | ✓ 1320ms | http |
| 103.165.138.173:8181 | ✓ 1705ms | 否 | ✓ 1404ms | ✓ 1437ms | ✓ 1141ms | http |
| 103.166.182.144:3128 | ✓ 1599ms | ✓ 1744ms | ✓ 1361ms | ✓ 1444ms | ✓ 1551ms | http |
| 61.52.131.172:8443 | ✓ 1058ms | ✓ 1352ms | ✓ 1678ms | ✓ 1486ms | ✓ 1194ms | http |
| 190.12.150.244:999 | ✓ 1534ms | ✓ 1705ms | ✓ 1089ms | 否 | 否 | http |
| 103.172.70.173:8080 | ✓ 1512ms | 否 | ✓ 1709ms | ✓ 1736ms | 否 | http |

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
