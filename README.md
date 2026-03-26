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

最后更新：2026-03-26 12:52:53 UTC（2026-03-26 20:52:53 UTC+8）

**代理总数：87**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 86 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 87 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.239.240:8800 | ✓ 1117ms | ✓ 1437ms | ✓ 765ms | ✓ 1468ms | ✓ 1332ms | http |
| 219.117.204.211:7799 | ✓ 1782ms | 否 | 否 | ✓ 1126ms | ✓ 1002ms | http |
| 167.103.115.102:8800 | ✓ 1745ms | 否 | ✓ 1169ms | ✓ 1364ms | ✓ 1212ms | http |
| 147.161.210.140:8800 | ✓ 1776ms | 否 | ✓ 1054ms | ✓ 1367ms | 否 | http |
| 5.104.87.17:8051 | ✓ 1940ms | 否 | ✓ 1354ms | ✓ 1366ms | 否 | http |
| 167.103.34.108:8800 | ✓ 1760ms | 否 | 否 | ✓ 1770ms | ✓ 1822ms | http |
| 35.225.22.61:80 | ✓ 554ms | 否 | ✓ 775ms | ✓ 1290ms | ✓ 868ms | http |
| 38.145.208.206:8448 | ✓ 492ms | ✓ 1380ms | ✓ 1480ms | ✓ 1165ms | ✓ 741ms | http |
| 194.67.99.223:1080 | ✓ 623ms | 否 | 否 | ✓ 1753ms | ✓ 1352ms | http |
| 113.160.132.26:8080 | ✓ 1647ms | ✓ 1664ms | ✓ 1574ms | ✓ 1517ms | ✓ 1184ms | http |
| 167.103.144.127:8800 | 否 | 否 | ✓ 1447ms | ✓ 1568ms | ✓ 1462ms | http |
| 116.80.65.78:3172 | ✓ 1650ms | 否 | ✓ 1635ms | 否 | ✓ 1796ms | http |
| 167.103.31.122:8800 | ✓ 1734ms | 否 | ✓ 1207ms | ✓ 1550ms | ✓ 1457ms | http |
| 178.218.105.99:3128 | 否 | 否 | ✓ 1643ms | ✓ 1788ms | ✓ 1452ms | http |
| 186.148.180.46:999 | ✓ 1413ms | 否 | ✓ 688ms | ✓ 1614ms | ✓ 1696ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1468ms | 否 | ✓ 1512ms | ✓ 1187ms | http |
| 160.250.4.13:1 | ✓ 1802ms | 否 | 否 | ✓ 1575ms | ✓ 1426ms | http |
| 38.145.220.60:8445 | ✓ 946ms | 否 | ✓ 622ms | ✓ 933ms | ✓ 1142ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1748ms | ✓ 1477ms | ✓ 1510ms | ✓ 1168ms | http |
| 137.184.1.87:3128 | ✓ 575ms | ✓ 1562ms | ✓ 1086ms | ✓ 1095ms | ✓ 924ms | http |
| 89.208.106.138:10808 | 否 | ✓ 1379ms | ✓ 626ms | ✓ 1393ms | 否 | http |
| 101.43.127.100:8877 | ✓ 1042ms | ✓ 1249ms | ✓ 1340ms | 否 | ✓ 1377ms | http |
| 24.199.124.152:3128 | ✓ 793ms | ✓ 1945ms | ✓ 1031ms | 否 | 否 | http |
| 51.79.71.106:8080 | ✓ 891ms | 否 | 否 | ✓ 1706ms | ✓ 1812ms | http |
| 59.46.216.131:30001 | ✓ 1133ms | 否 | ✓ 1860ms | 否 | ✓ 1180ms | http |
| 45.140.147.82:1081 | ✓ 696ms | 否 | ✓ 1206ms | ✓ 1255ms | ✓ 1035ms | http |
| 38.34.179.67:8453 | ✓ 1123ms | 否 | ✓ 953ms | ✓ 1142ms | 否 | http |
| 45.144.28.81:10808 | ✓ 1162ms | 否 | ✓ 1239ms | ✓ 1204ms | ✓ 1440ms | http |
| 115.231.181.40:8128 | ✓ 1471ms | ✓ 1315ms | ✓ 1158ms | ✓ 1413ms | 否 | http |
| 172.233.153.101:3128 | ✓ 691ms | 否 | ✓ 1815ms | ✓ 897ms | ✓ 752ms | http |
| 103.84.95.54:7890 | ✓ 1020ms | 否 | ✓ 920ms | ✓ 1130ms | 否 | http |
| 38.145.208.237:8445 | ✓ 857ms | 否 | ✓ 1288ms | ✓ 1111ms | ✓ 899ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1033ms | ✓ 1494ms | ✓ 1167ms | http |
| 38.34.179.51:8449 | ✓ 368ms | ✓ 1089ms | 否 | ✓ 1562ms | ✓ 1144ms | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1803ms | ✓ 1778ms | ✓ 1392ms | http |
| 20.118.221.52:3128 | ✓ 1434ms | 否 | ✓ 1319ms | ✓ 1170ms | ✓ 923ms | http |
| 115.231.37.166:14286 | ✓ 1196ms | ✓ 1946ms | ✓ 1340ms | 否 | ✓ 1834ms | http |
| 181.41.201.85:3128 | ✓ 1549ms | 否 | ✓ 1728ms | 否 | ✓ 1761ms | http |
| 115.231.37.247:14729 | ✓ 1672ms | 否 | ✓ 1258ms | ✓ 1582ms | ✓ 1569ms | http |
| 114.237.77.220:1080 | ✓ 1153ms | 否 | 否 | ✓ 1426ms | ✓ 1124ms | http |
| 139.5.189.229:8888 | ✓ 1720ms | 否 | ✓ 1690ms | 否 | ✓ 1649ms | http |
| 38.145.208.227:8452 | ✓ 630ms | ✓ 1732ms | ✓ 579ms | ✓ 1338ms | ✓ 779ms | http |
| 38.34.179.71:8448 | ✓ 688ms | 否 | ✓ 301ms | ✓ 995ms | ✓ 1377ms | http |
| 45.136.130.195:8446 | ✓ 514ms | ✓ 1284ms | ✓ 1527ms | ✓ 1146ms | ✓ 983ms | http |
| 183.249.5.111:22222 | ✓ 935ms | ✓ 1194ms | ✓ 889ms | ✓ 1393ms | ✓ 889ms | http |
| 38.145.208.204:8451 | ✓ 600ms | 否 | ✓ 1191ms | ✓ 1359ms | ✓ 1365ms | http |
| 38.145.220.6:8446 | ✓ 736ms | ✓ 902ms | ✓ 795ms | 否 | ✓ 728ms | http |
| 38.34.179.193:8451 | ✓ 409ms | 否 | ✓ 1789ms | ✓ 921ms | ✓ 1074ms | http |
| 45.136.130.247:8450 | ✓ 1171ms | 否 | ✓ 982ms | ✓ 1021ms | ✓ 1528ms | http |
| 38.145.208.169:8453 | ✓ 402ms | ✓ 1124ms | 否 | ✓ 1281ms | ✓ 690ms | http |
| 38.145.218.227:8450 | ✓ 415ms | ✓ 1126ms | ✓ 1841ms | ✓ 1500ms | ✓ 708ms | http |
| 38.145.218.87:8450 | ✓ 404ms | ✓ 1706ms | 否 | ✓ 900ms | ✓ 883ms | http |
| 38.145.208.189:8450 | ✓ 632ms | 否 | ✓ 1886ms | ✓ 1072ms | ✓ 834ms | http |
| 38.145.218.14:8450 | ✓ 1491ms | ✓ 1495ms | ✓ 960ms | 否 | ✓ 1045ms | http |
| 45.136.130.180:8450 | ✓ 1579ms | 否 | ✓ 1173ms | ✓ 1880ms | ✓ 1083ms | http |
| 38.34.179.178:8444 | ✓ 1573ms | ✓ 1893ms | ✓ 1360ms | ✓ 1471ms | ✓ 1595ms | http |
| 38.34.179.201:8451 | ✓ 1959ms | ✓ 1397ms | ✓ 1202ms | 否 | ✓ 801ms | http |
| 38.34.179.39:8445 | ✓ 1459ms | ✓ 991ms | ✓ 1504ms | 否 | ✓ 1119ms | http |
| 38.145.203.43:8451 | ✓ 1714ms | 否 | ✓ 944ms | ✓ 1473ms | 否 | http |
| 45.136.131.63:8453 | ✓ 1772ms | ✓ 1446ms | ✓ 1010ms | 否 | ✓ 798ms | http |
| 158.160.215.167:8124 | ✓ 713ms | 否 | ✓ 646ms | 否 | ✓ 1404ms | http |
| 38.34.179.152:8450 | 否 | 否 | ✓ 1227ms | ✓ 1918ms | ✓ 1011ms | http |
| 158.160.215.167:8127 | 否 | 否 | ✓ 1569ms | ✓ 1682ms | ✓ 1599ms | http |
| 38.145.203.98:8445 | ✓ 1843ms | ✓ 1220ms | ✓ 1557ms | ✓ 1696ms | ✓ 1235ms | http |
| 45.136.131.32:8445 | 否 | ✓ 1784ms | ✓ 357ms | ✓ 1024ms | ✓ 714ms | http |
| 222.184.48.251:22222 | ✓ 1086ms | ✓ 1316ms | ✓ 1480ms | 否 | 否 | http |
| 183.249.5.110:22222 | ✓ 934ms | ✓ 1289ms | ✓ 956ms | ✓ 1162ms | ✓ 885ms | http |
| 183.249.5.117:22222 | ✓ 928ms | ✓ 1124ms | ✓ 991ms | ✓ 1217ms | ✓ 1024ms | http |
| 183.249.5.105:22222 | ✓ 1026ms | ✓ 1392ms | ✓ 940ms | ✓ 1472ms | ✓ 905ms | http |
| 180.127.230.46:1080 | ✓ 1070ms | ✓ 1450ms | ✓ 1190ms | 否 | ✓ 1980ms | http |
| 113.255.59.226:8080 | ✓ 1192ms | 否 | 否 | ✓ 1358ms | ✓ 1427ms | http |
| 45.186.6.104:3128 | ✓ 953ms | ✓ 1554ms | ✓ 1786ms | 否 | 否 | http |
| 38.34.179.16:8451 | 否 | ✓ 1779ms | 否 | ✓ 1153ms | ✓ 1627ms | http |
| 115.231.37.227:13914 | 否 | ✓ 1753ms | ✓ 1402ms | ✓ 1859ms | 否 | http |
| 88.80.150.82:8080 | ✓ 858ms | 否 | ✓ 1742ms | ✓ 1856ms | ✓ 1505ms | https |
| 223.16.170.103:80 | ✓ 1693ms | 否 | ✓ 1267ms | ✓ 1314ms | ✓ 1358ms | http |
| 106.75.15.167:7890 | ✓ 1729ms | 否 | ✓ 1923ms | ✓ 1976ms | 否 | http |
| 91.233.223.147:3128 | ✓ 1733ms | 否 | ✓ 1568ms | 否 | ✓ 1964ms | http |
| 103.113.70.189:1081 | ✓ 321ms | 否 | ✓ 850ms | ✓ 985ms | ✓ 737ms | http |
| 16.78.119.130:443 | 否 | ✓ 1990ms | ✓ 1717ms | 否 | ✓ 1865ms | http |
| 85.208.108.43:10808 | ✓ 1219ms | 否 | ✓ 217ms | 否 | ✓ 916ms | http |
| 77.110.113.24:40000 | ✓ 1072ms | 否 | ✓ 1372ms | 否 | ✓ 1981ms | http |
| 85.208.108.43:2094 | ✓ 397ms | 否 | ✓ 612ms | 否 | ✓ 1785ms | http |
| 5.102.109.41:999 | ✓ 384ms | ✓ 1377ms | ✓ 310ms | ✓ 1521ms | ✓ 1044ms | http |
| 222.184.48.242:22222 | ✓ 1143ms | ✓ 1435ms | ✓ 1157ms | 否 | ✓ 1952ms | http |
| 202.141.161.53:30001 | 否 | ✓ 1606ms | ✓ 1421ms | 否 | ✓ 1775ms | http |
| 38.34.179.88:8446 | ✓ 989ms | ✓ 1681ms | ✓ 1320ms | ✓ 955ms | 否 | http |

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
