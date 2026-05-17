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

最后更新：2026-05-17 22:48:32 UTC（2026-05-18 06:48:32 UTC+8）

**代理总数：86**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 86 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 86 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 51.161.50.166:3128 | ✓ 554ms | ✓ 1499ms | ✓ 1131ms | ✓ 1077ms | ✓ 1063ms | http |
| 170.106.136.181:31002 | ✓ 841ms | 否 | 否 | ✓ 1175ms | ✓ 1228ms | http |
| 185.200.188.234:10001 | ✓ 1503ms | 否 | ✓ 1633ms | 否 | ✓ 1555ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1819ms | ✓ 1164ms | ✓ 1451ms | ✓ 1212ms | http |
| 212.58.132.5:8888 | ✓ 1858ms | 否 | ✓ 1593ms | ✓ 1543ms | ✓ 1256ms | http |
| 115.231.181.40:8128 | ✓ 1128ms | 否 | ✓ 1248ms | 否 | ✓ 1722ms | http |
| 45.125.67.37:8443 | ✓ 1165ms | 否 | ✓ 1061ms | ✓ 1279ms | 否 | http |
| 218.108.131.186:17890 | 否 | ✓ 1279ms | ✓ 1280ms | ✓ 1328ms | ✓ 1276ms | http |
| 20.164.75.153:8080 | ✓ 1333ms | 否 | ✓ 1163ms | 否 | ✓ 1853ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1311ms | ✓ 1937ms | ✓ 1114ms | http |
| 5.252.33.13:2025 | ✓ 1714ms | ✓ 1962ms | ✓ 1234ms | 否 | ✓ 1710ms | http |
| 168.222.254.136:8888 | ✓ 1026ms | ✓ 1792ms | 否 | 否 | ✓ 1707ms | http |
| 84.47.150.125:1080 | ✓ 918ms | 否 | ✓ 1152ms | ✓ 1998ms | ✓ 1843ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1856ms | 否 | ✓ 1622ms | ✓ 1307ms | http |
| 158.160.215.167:8124 | ✓ 920ms | ✓ 1924ms | ✓ 1452ms | 否 | ✓ 1434ms | http |
| 8.154.21.175:3128 | ✓ 1037ms | ✓ 1242ms | ✓ 1103ms | ✓ 1387ms | ✓ 1150ms | http |
| 128.199.113.85:9090 | ✓ 1406ms | 否 | ✓ 1098ms | ✓ 1860ms | ✓ 1213ms | http |
| 120.92.212.16:7890 | ✓ 1920ms | ✓ 1845ms | ✓ 1177ms | ✓ 1342ms | ✓ 1171ms | http |
| 120.92.212.16:8890 | ✓ 1556ms | ✓ 1492ms | ✓ 1424ms | 否 | 否 | http |
| 47.105.98.23:3128 | ✓ 1325ms | 否 | ✓ 1599ms | ✓ 1451ms | ✓ 1357ms | http |
| 43.156.90.221:10808 | ✓ 940ms | 否 | 否 | ✓ 1276ms | ✓ 1442ms | http |
| 128.199.254.13:9090 | 否 | 否 | ✓ 1025ms | ✓ 1529ms | ✓ 1132ms | http |
| 148.230.4.241:999 | 否 | ✓ 1412ms | ✓ 915ms | ✓ 1494ms | ✓ 1191ms | http |
| 57.129.144.178:40000 | ✓ 1016ms | 否 | ✓ 1490ms | ✓ 1668ms | ✓ 1432ms | http |
| 103.147.152.12:1080 | ✓ 431ms | 否 | ✓ 1953ms | ✓ 1545ms | ✓ 1145ms | http |
| 178.156.224.42:3128 | ✓ 1935ms | 否 | ✓ 1881ms | 否 | ✓ 1627ms | http |
| 103.147.152.12:1095 | 否 | 否 | ✓ 765ms | ✓ 1493ms | ✓ 1159ms | http |
| 38.180.2.107:3128 | ✓ 845ms | ✓ 1647ms | ✓ 1915ms | 否 | ✓ 1830ms | http |
| 129.80.238.83:444 | ✓ 182ms | 否 | ✓ 1362ms | ✓ 931ms | ✓ 687ms | http |
| 129.80.217.21:444 | ✓ 184ms | 否 | 否 | ✓ 896ms | ✓ 669ms | http |
| 91.242.229.129:8092 | ✓ 1863ms | ✓ 1314ms | ✓ 1324ms | 否 | ✓ 1340ms | http |
| 210.223.44.230:3128 | ✓ 1395ms | ✓ 1590ms | 否 | ✓ 1478ms | 否 | http |
| 180.125.216.109:8118 | ✓ 1160ms | ✓ 1383ms | 否 | 否 | ✓ 1138ms | http |
| 61.155.242.150:5566 | ✓ 1293ms | ✓ 1366ms | ✓ 1499ms | ✓ 1471ms | ✓ 1372ms | http |
| 64.188.77.221:3128 | ✓ 619ms | 否 | ✓ 800ms | 否 | ✓ 1343ms | http |
| 64.188.77.26:3128 | ✓ 602ms | 否 | 否 | ✓ 1639ms | ✓ 1264ms | http |
| 42.200.76.16:3888 | 否 | 否 | ✓ 1395ms | ✓ 1084ms | ✓ 869ms | http |
| 82.65.27.56:80 | ✓ 1561ms | 否 | ✓ 1574ms | ✓ 1539ms | 否 | http |
| 192.210.140.36:7913 | ✓ 1926ms | ✓ 1036ms | ✓ 1062ms | ✓ 1254ms | ✓ 855ms | http |
| 166.88.55.83:7890 | ✓ 808ms | ✓ 1300ms | ✓ 799ms | ✓ 1007ms | ✓ 806ms | http |
| 1.231.81.166:3128 | ✓ 1835ms | ✓ 1240ms | 否 | ✓ 1201ms | ✓ 973ms | http |
| 121.230.9.113:1080 | ✓ 1384ms | ✓ 1566ms | ✓ 1365ms | ✓ 1823ms | ✓ 1321ms | http |
| 113.176.92.71:3128 | ✓ 1823ms | ✓ 1516ms | ✓ 1361ms | ✓ 1588ms | ✓ 1510ms | http |
| 62.133.60.5:3128 | 否 | 否 | ✓ 701ms | ✓ 1729ms | ✓ 1202ms | http |
| 141.148.170.136:3128 | ✓ 1897ms | ✓ 1798ms | 否 | 否 | ✓ 1588ms | http |
| 46.105.190.38:3128 | ✓ 445ms | ✓ 1267ms | ✓ 701ms | ✓ 1673ms | ✓ 1285ms | http |
| 104.248.151.93:9090 | ✓ 950ms | 否 | ✓ 925ms | ✓ 1345ms | 否 | http |
| 128.199.116.219:9090 | ✓ 1055ms | 否 | ✓ 1036ms | ✓ 1543ms | 否 | http |
| 159.223.41.216:9090 | ✓ 942ms | 否 | ✓ 942ms | ✓ 1310ms | ✓ 1094ms | http |
| 121.130.177.28:8888 | ✓ 1315ms | 否 | ✓ 1987ms | ✓ 1982ms | ✓ 1551ms | http |
| 14.29.168.215:1080 | ✓ 1409ms | ✓ 1411ms | 否 | ✓ 1341ms | ✓ 1073ms | http |
| 101.32.243.189:80 | ✓ 1386ms | 否 | ✓ 1562ms | 否 | ✓ 1962ms | http |
| 185.21.15.206:3128 | ✓ 1337ms | 否 | ✓ 649ms | ✓ 1831ms | ✓ 1428ms | http |
| 3.101.133.120:80 | ✓ 507ms | ✓ 1315ms | ✓ 1580ms | ✓ 1276ms | ✓ 1047ms | http |
| 138.2.239.213:10010 | ✓ 1004ms | ✓ 1996ms | 否 | ✓ 1388ms | ✓ 1100ms | http |
| 158.160.215.167:8123 | ✓ 1090ms | 否 | ✓ 979ms | ✓ 1865ms | 否 | http |
| 152.42.170.187:9090 | 否 | 否 | ✓ 1470ms | ✓ 1506ms | ✓ 1146ms | http |
| 128.199.121.61:9090 | ✓ 1384ms | 否 | ✓ 1110ms | ✓ 1528ms | ✓ 1011ms | http |
| 109.234.38.35:3128 | ✓ 1360ms | ✓ 1685ms | ✓ 1990ms | 否 | ✓ 1531ms | http |
| 168.138.202.218:3128 | ✓ 1667ms | ✓ 1645ms | ✓ 1625ms | ✓ 1219ms | ✓ 1094ms | http |
| 167.71.196.178:80 | ✓ 1080ms | 否 | ✓ 1212ms | ✓ 1290ms | ✓ 1050ms | http |
| 94.247.244.120:3128 | ✓ 1023ms | ✓ 1575ms | ✓ 1124ms | 否 | 否 | http |
| 103.82.23.118:5226 | ✓ 1688ms | 否 | ✓ 1539ms | 否 | ✓ 1930ms | http |
| 5.129.223.50:3128 | ✓ 1022ms | ✓ 1765ms | ✓ 668ms | ✓ 1884ms | ✓ 1652ms | http |
| 45.129.141.143:3128 | ✓ 1507ms | ✓ 1900ms | ✓ 1617ms | 否 | ✓ 1659ms | http |
| 45.186.6.104:3128 | ✓ 1816ms | ✓ 1788ms | ✓ 1660ms | 否 | 否 | http |
| 103.134.85.145:3128 | ✓ 1064ms | 否 | ✓ 1063ms | ✓ 1512ms | ✓ 1185ms | http |
| 139.135.170.12:8082 | 否 | 否 | ✓ 1542ms | ✓ 1627ms | ✓ 1609ms | http |
| 103.239.201.50:1 | 否 | 否 | ✓ 1501ms | ✓ 1612ms | ✓ 1709ms | http |
| 146.190.80.158:9090 | 否 | 否 | ✓ 1658ms | ✓ 1603ms | ✓ 1035ms | http |
| 8.219.97.248:80 | 否 | 否 | ✓ 1404ms | ✓ 1767ms | ✓ 1490ms | http |
| 2.27.32.81:3128 | ✓ 755ms | 否 | ✓ 721ms | ✓ 1896ms | ✓ 1714ms | http |
| 178.63.155.151:8898 | ✓ 1247ms | 否 | ✓ 647ms | ✓ 1641ms | 否 | http |
| 152.32.132.190:7890 | ✓ 1109ms | ✓ 1922ms | ✓ 1064ms | ✓ 1498ms | ✓ 1506ms | http |
| 114.214.165.78:10810 | ✓ 1678ms | ✓ 1998ms | ✓ 1930ms | ✓ 1664ms | ✓ 1620ms | http |
| 94.131.122.128:1081 | 否 | 否 | ✓ 893ms | ✓ 1247ms | ✓ 878ms | http |
| 113.45.216.128:55555 | ✓ 1263ms | ✓ 1534ms | 否 | ✓ 1568ms | ✓ 1268ms | http |
| 114.214.170.41:27890 | ✓ 1320ms | ✓ 1617ms | ✓ 1539ms | ✓ 1635ms | ✓ 1277ms | http |
| 34.96.238.40:8080 | ✓ 1433ms | ✓ 1313ms | 否 | ✓ 1674ms | 否 | http |
| 128.199.114.189:9090 | ✓ 1383ms | 否 | 否 | ✓ 1831ms | ✓ 1506ms | http |
| 103.19.228.4:8080 | ✓ 1718ms | 否 | 否 | ✓ 1882ms | ✓ 1855ms | http |
| 103.244.107.150:8080 | 否 | 否 | ✓ 1994ms | ✓ 1694ms | ✓ 1646ms | http |
| 103.172.70.173:8080 | ✓ 1578ms | 否 | ✓ 1899ms | ✓ 1686ms | ✓ 1678ms | http |
| 190.14.224.244:999 | 否 | 否 | ✓ 1681ms | ✓ 1802ms | ✓ 1474ms | http |
| 137.59.47.73:3128 | 否 | ✓ 1592ms | ✓ 1162ms | 否 | ✓ 1824ms | http |
| 212.34.146.118:3128 | ✓ 707ms | 否 | ✓ 864ms | ✓ 1537ms | ✓ 1730ms | http |

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
