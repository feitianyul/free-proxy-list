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

最后更新：2026-05-16 14:28:36 UTC（2026-05-16 22:28:36 UTC+8）

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
| 107.175.85.198:1080 | ✓ 523ms | 否 | ✓ 902ms | ✓ 1413ms | ✓ 1182ms | http |
| 42.96.16.158:1311 | ✓ 1824ms | 否 | ✓ 1211ms | ✓ 1181ms | ✓ 974ms | http |
| 113.160.132.26:8080 | ✓ 1826ms | ✓ 1809ms | ✓ 1186ms | ✓ 1227ms | ✓ 1048ms | http |
| 185.200.188.234:10001 | ✓ 1558ms | 否 | ✓ 1796ms | 否 | ✓ 1856ms | http |
| 192.210.140.36:7913 | 否 | ✓ 1058ms | ✓ 487ms | ✓ 1408ms | ✓ 1373ms | http |
| 45.125.67.37:8443 | 否 | 否 | ✓ 1210ms | ✓ 1162ms | ✓ 1636ms | http |
| 218.108.131.186:17890 | ✓ 1568ms | ✓ 1172ms | ✓ 1013ms | 否 | 否 | http |
| 8.219.97.248:80 | ✓ 1358ms | 否 | ✓ 1324ms | ✓ 1356ms | 否 | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1054ms | ✓ 1142ms | ✓ 901ms | http |
| 45.88.0.115:3128 | ✓ 1705ms | ✓ 1558ms | 否 | 否 | ✓ 1105ms | http |
| 45.88.0.98:3128 | ✓ 1444ms | 否 | 否 | ✓ 1446ms | ✓ 1242ms | http |
| 199.38.85.122:40014 | ✓ 340ms | 否 | ✓ 314ms | ✓ 1415ms | ✓ 1225ms | http |
| 20.210.76.175:8561 | 否 | ✓ 1604ms | ✓ 1087ms | ✓ 1592ms | ✓ 1441ms | http |
| 20.210.76.104:8561 | 否 | ✓ 1719ms | ✓ 1065ms | ✓ 1505ms | ✓ 1447ms | http |
| 20.18.193.135:8561 | 否 | ✓ 1735ms | ✓ 1049ms | ✓ 1505ms | ✓ 1444ms | http |
| 20.27.15.49:8561 | 否 | 否 | ✓ 1004ms | ✓ 1364ms | ✓ 1381ms | http |
| 20.210.76.178:8561 | 否 | 否 | ✓ 1013ms | ✓ 1366ms | ✓ 1369ms | http |
| 115.231.181.40:8128 | ✓ 983ms | ✓ 1733ms | ✓ 898ms | ✓ 1864ms | ✓ 1002ms | http |
| 103.21.220.141:3128 | ✓ 690ms | 否 | ✓ 691ms | ✓ 868ms | ✓ 710ms | http |
| 128.199.121.61:9090 | ✓ 1374ms | 否 | ✓ 1278ms | ✓ 1087ms | ✓ 909ms | http |
| 128.199.254.13:9090 | ✓ 1369ms | 否 | ✓ 1339ms | ✓ 1090ms | ✓ 851ms | http |
| 128.199.114.189:9090 | ✓ 1364ms | 否 | ✓ 1429ms | ✓ 1122ms | ✓ 895ms | http |
| 146.190.80.158:9090 | 否 | 否 | ✓ 951ms | ✓ 1086ms | ✓ 872ms | http |
| 152.42.170.187:9090 | 否 | 否 | ✓ 1933ms | ✓ 1267ms | ✓ 896ms | http |
| 148.230.4.241:999 | ✓ 1099ms | ✓ 1635ms | ✓ 759ms | ✓ 1413ms | ✓ 1361ms | http |
| 84.47.150.125:1080 | ✓ 1294ms | 否 | ✓ 1481ms | 否 | ✓ 1628ms | http |
| 170.106.136.181:31002 | ✓ 346ms | ✓ 619ms | ✓ 475ms | ✓ 698ms | ✓ 527ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1302ms | ✓ 1661ms | ✓ 1167ms | ✓ 1126ms | http |
| 43.156.90.221:10808 | ✓ 1651ms | 否 | ✓ 1461ms | 否 | ✓ 1221ms | http |
| 91.242.229.129:8092 | ✓ 1558ms | ✓ 1853ms | 否 | 否 | ✓ 1459ms | http |
| 8.154.21.175:3128 | ✓ 1893ms | ✓ 1049ms | ✓ 1331ms | ✓ 1952ms | ✓ 949ms | http |
| 103.157.200.126:3128 | ✓ 1646ms | 否 | ✓ 1235ms | ✓ 1891ms | ✓ 1474ms | http |
| 138.2.239.213:10010 | ✓ 1212ms | ✓ 1976ms | 否 | ✓ 1203ms | ✓ 1739ms | http |
| 5.129.248.58:3128 | ✓ 1229ms | ✓ 1684ms | ✓ 1532ms | ✓ 1722ms | ✓ 1754ms | http |
| 152.32.132.190:7890 | ✓ 998ms | ✓ 1270ms | ✓ 1540ms | ✓ 1674ms | 否 | http |
| 20.78.118.91:8561 | ✓ 1384ms | ✓ 1203ms | ✓ 523ms | ✓ 834ms | ✓ 756ms | http |
| 20.210.39.153:8561 | ✓ 1385ms | ✓ 1078ms | ✓ 616ms | ✓ 865ms | ✓ 764ms | http |
| 20.78.26.206:8561 | ✓ 1386ms | ✓ 1642ms | ✓ 521ms | ✓ 923ms | ✓ 757ms | http |
| 199.38.85.122:40010 | ✓ 1727ms | ✓ 1744ms | ✓ 1741ms | 否 | ✓ 1275ms | http |
| 20.27.15.111:8561 | ✓ 1385ms | ✓ 1659ms | 否 | ✓ 1823ms | ✓ 1488ms | http |
| 20.27.11.248:8561 | ✓ 1387ms | ✓ 1648ms | 否 | ✓ 1831ms | ✓ 1490ms | http |
| 20.27.13.35:8561 | ✓ 1386ms | ✓ 1662ms | 否 | ✓ 1823ms | ✓ 1494ms | http |
| 20.27.14.220:8561 | ✓ 1385ms | 否 | ✓ 1971ms | ✓ 1700ms | ✓ 1446ms | http |
| 45.88.0.99:3128 | 否 | 否 | ✓ 593ms | ✓ 1440ms | ✓ 1111ms | http |
| 166.88.55.83:7890 | ✓ 714ms | ✓ 1180ms | ✓ 742ms | ✓ 888ms | ✓ 712ms | http |
| 49.144.24.53:8082 | ✓ 1390ms | 否 | 否 | ✓ 1663ms | ✓ 1484ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1267ms | ✓ 1314ms | 否 | ✓ 801ms | http |
| 128.199.116.219:9090 | 否 | 否 | ✓ 747ms | ✓ 1077ms | ✓ 867ms | http |
| 34.71.229.255:3128 | ✓ 1709ms | 否 | ✓ 1618ms | ✓ 1487ms | ✓ 1123ms | http |
| 43.167.192.85:8080 | ✓ 581ms | ✓ 1078ms | ✓ 522ms | ✓ 832ms | ✓ 652ms | http |
| 128.199.113.85:9090 | ✓ 922ms | 否 | ✓ 765ms | ✓ 1443ms | ✓ 940ms | http |
| 159.89.31.62:8080 | ✓ 661ms | ✓ 1950ms | ✓ 1903ms | ✓ 1703ms | ✓ 1870ms | http |
| 217.174.244.117:3129 | ✓ 981ms | 否 | ✓ 1761ms | 否 | ✓ 1994ms | http |
| 3.101.133.120:80 | ✓ 439ms | 否 | ✓ 1066ms | ✓ 1159ms | ✓ 775ms | http |
| 45.88.0.114:3128 | ✓ 639ms | 否 | ✓ 940ms | ✓ 1809ms | 否 | http |
| 103.242.105.76:8090 | ✓ 1270ms | 否 | ✓ 1575ms | ✓ 1503ms | ✓ 1600ms | http |
| 14.190.190.211:20114 | ✓ 1248ms | 否 | ✓ 1904ms | ✓ 1417ms | ✓ 1630ms | http |
| 137.59.47.73:3128 | 否 | 否 | ✓ 1045ms | ✓ 1145ms | ✓ 941ms | http |
| 114.214.170.41:27890 | ✓ 1200ms | ✓ 1471ms | ✓ 1362ms | ✓ 1440ms | 否 | http |
| 103.147.152.12:1095 | ✓ 1122ms | 否 | ✓ 1214ms | ✓ 1670ms | ✓ 1294ms | http |
| 103.147.152.12:1080 | ✓ 562ms | 否 | ✓ 967ms | ✓ 1652ms | ✓ 1297ms | http |
| 45.88.0.117:3128 | ✓ 611ms | ✓ 1443ms | ✓ 935ms | 否 | 否 | http |
| 113.45.216.128:55555 | ✓ 1020ms | ✓ 1357ms | ✓ 1132ms | ✓ 1386ms | ✓ 1101ms | http |
| 212.58.132.5:8888 | 否 | 否 | ✓ 1441ms | ✓ 1574ms | ✓ 1272ms | http |
| 59.46.216.131:30001 | ✓ 1088ms | 否 | ✓ 1105ms | 否 | ✓ 1149ms | http |
| 104.248.151.93:9090 | ✓ 1121ms | 否 | ✓ 743ms | ✓ 1082ms | ✓ 859ms | http |
| 116.171.106.26:3443 | 否 | ✓ 1493ms | ✓ 1669ms | ✓ 1709ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1769ms | 否 | 否 | ✓ 1452ms | ✓ 1316ms | http |
| 5.75.139.30:1081 | ✓ 807ms | 否 | ✓ 1774ms | 否 | ✓ 1679ms | http |
| 38.188.247.12:999 | ✓ 833ms | ✓ 1545ms | ✓ 1155ms | 否 | 否 | http |
| 160.238.65.7:3128 | ✓ 1743ms | 否 | 否 | ✓ 1445ms | ✓ 1971ms | http |
| 160.238.65.2:3128 | ✓ 1770ms | 否 | ✓ 1493ms | 否 | ✓ 1550ms | http |
| 103.69.84.106:3131 | ✓ 1699ms | 否 | ✓ 1207ms | ✓ 1176ms | ✓ 943ms | http |
| 61.52.131.172:8443 | ✓ 886ms | ✓ 1205ms | ✓ 1352ms | ✓ 1215ms | ✓ 1982ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1161ms | ✓ 1660ms | ✓ 1285ms | ✓ 1996ms | http |
| 103.157.117.116:8080 | 否 | 否 | ✓ 1706ms | ✓ 1794ms | ✓ 1836ms | http |

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
