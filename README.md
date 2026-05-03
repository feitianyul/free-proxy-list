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

最后更新：2026-05-03 19:57:54 UTC（2026-05-04 03:57:54 UTC+8）

**代理总数：70**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 70 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 70 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 218.108.131.186:17890 | ✓ 981ms | ✓ 1290ms | ✓ 1016ms | ✓ 1308ms | ✓ 1071ms | http |
| 1.231.81.166:3128 | ✓ 1832ms | 否 | 否 | ✓ 1291ms | ✓ 1208ms | http |
| 113.160.132.26:8080 | ✓ 1693ms | ✓ 1793ms | 否 | ✓ 1454ms | ✓ 1459ms | http |
| 45.167.124.71:999 | ✓ 1542ms | 否 | ✓ 1306ms | ✓ 1977ms | ✓ 1579ms | http |
| 62.113.119.14:8080 | 否 | 否 | ✓ 1063ms | ✓ 1483ms | ✓ 1153ms | http |
| 86.104.74.110:1081 | ✓ 554ms | ✓ 1414ms | ✓ 600ms | ✓ 1582ms | ✓ 1304ms | http |
| 154.64.232.35:8080 | ✓ 1479ms | 否 | ✓ 1076ms | ✓ 1106ms | ✓ 783ms | http |
| 206.206.126.177:2412 | ✓ 1077ms | 否 | ✓ 1572ms | ✓ 1210ms | ✓ 951ms | http |
| 37.187.109.70:10111 | ✓ 1015ms | ✓ 1560ms | ✓ 1029ms | 否 | ✓ 1259ms | http |
| 46.105.190.40:3128 | 否 | ✓ 1917ms | ✓ 815ms | ✓ 1703ms | 否 | http |
| 152.32.132.190:7890 | ✓ 1849ms | 否 | ✓ 1226ms | 否 | ✓ 1842ms | http |
| 47.77.216.82:1080 | 否 | 否 | ✓ 1762ms | ✓ 1135ms | ✓ 1013ms | http |
| 86.104.72.220:1082 | ✓ 635ms | ✓ 1254ms | ✓ 1082ms | ✓ 1215ms | ✓ 1864ms | http |
| 80.92.204.47:1081 | ✓ 896ms | 否 | ✓ 466ms | ✓ 1553ms | 否 | http |
| 77.110.107.80:8080 | ✓ 384ms | ✓ 1377ms | ✓ 970ms | ✓ 1666ms | ✓ 1384ms | http |
| 212.58.132.5:8888 | ✓ 1010ms | 否 | ✓ 1035ms | ✓ 1423ms | ✓ 1116ms | http |
| 47.85.51.197:1080 | ✓ 299ms | 否 | ✓ 563ms | ✓ 1050ms | ✓ 833ms | http |
| 46.105.190.38:3128 | ✓ 1162ms | ✓ 1272ms | ✓ 918ms | 否 | 否 | http |
| 15.204.238.117:3128 | ✓ 215ms | ✓ 1796ms | ✓ 504ms | 否 | 否 | http |
| 86.104.72.220:1081 | ✓ 240ms | ✓ 875ms | ✓ 336ms | ✓ 1093ms | 否 | http |
| 152.70.91.193:40000 | ✓ 1987ms | 否 | 否 | ✓ 1638ms | ✓ 1623ms | http |
| 147.45.178.211:14658 | 否 | ✓ 1669ms | ✓ 1301ms | ✓ 1968ms | ✓ 1612ms | http |
| 91.233.223.147:3128 | ✓ 1510ms | ✓ 1864ms | ✓ 1229ms | 否 | ✓ 1524ms | http |
| 46.39.105.157:8080 | ✓ 657ms | 否 | ✓ 1378ms | 否 | ✓ 1337ms | http |
| 103.157.200.126:3128 | ✓ 1028ms | 否 | ✓ 1241ms | 否 | ✓ 1307ms | http |
| 62.60.149.161:3128 | ✓ 1097ms | ✓ 1644ms | ✓ 1624ms | 否 | 否 | http |
| 43.99.54.236:5555 | ✓ 818ms | ✓ 1123ms | ✓ 906ms | ✓ 1017ms | ✓ 816ms | http |
| 86.104.74.110:1082 | ✓ 995ms | 否 | ✓ 1167ms | ✓ 1589ms | ✓ 1117ms | http |
| 94.131.118.129:1081 | ✓ 1609ms | ✓ 1131ms | ✓ 1607ms | ✓ 1512ms | ✓ 1324ms | http |
| 94.131.118.129:1082 | ✓ 1827ms | ✓ 1197ms | ✓ 836ms | 否 | ✓ 1378ms | http |
| 45.140.147.82:1081 | ✓ 911ms | 否 | ✓ 1266ms | ✓ 1706ms | 否 | http |
| 109.120.156.122:8090 | ✓ 1241ms | ✓ 1945ms | ✓ 653ms | 否 | ✓ 1613ms | http |
| 104.128.138.186:1080 | ✓ 1246ms | 否 | 否 | ✓ 1856ms | ✓ 1711ms | http |
| 157.230.220.25:4857 | ✓ 179ms | ✓ 1424ms | 否 | ✓ 893ms | 否 | http |
| 81.26.190.143:1080 | ✓ 1983ms | ✓ 1937ms | ✓ 613ms | ✓ 1767ms | ✓ 1459ms | http |
| 8.219.97.248:80 | ✓ 1564ms | 否 | ✓ 1492ms | 否 | ✓ 1765ms | http |
| 8.154.21.175:3128 | ✓ 1035ms | ✓ 1268ms | ✓ 995ms | ✓ 1267ms | ✓ 1091ms | http |
| 168.110.52.228:3128 | 否 | ✓ 1143ms | 否 | ✓ 1091ms | ✓ 844ms | http |
| 45.125.67.37:8443 | ✓ 959ms | 否 | ✓ 1045ms | ✓ 1313ms | ✓ 1212ms | http |
| 135.125.232.193:3128 | ✓ 1174ms | ✓ 1638ms | 否 | 否 | ✓ 1752ms | http |
| 103.3.246.71:3128 | ✓ 1830ms | 否 | ✓ 1514ms | ✓ 1766ms | ✓ 1354ms | http |
| 51.145.178.158:3128 | ✓ 1147ms | ✓ 1735ms | 否 | 否 | ✓ 1594ms | http |
| 172.208.25.199:3128 | ✓ 238ms | ✓ 1586ms | ✓ 798ms | ✓ 1996ms | ✓ 1453ms | http |
| 59.19.47.77:3100 | 否 | ✓ 1718ms | ✓ 1713ms | ✓ 1610ms | 否 | http |
| 157.20.207.67:1111 | 否 | 否 | ✓ 1929ms | ✓ 1737ms | ✓ 1966ms | http |
| 103.93.93.221:8181 | 否 | 否 | ✓ 1815ms | ✓ 1821ms | ✓ 1753ms | http |
| 201.71.24.65:8082 | ✓ 984ms | ✓ 1998ms | ✓ 1875ms | 否 | ✓ 1862ms | http |
| 49.147.116.20:8082 | ✓ 1745ms | 否 | ✓ 1856ms | ✓ 1697ms | 否 | http |
| 193.123.250.39:1080 | ✓ 1520ms | 否 | 否 | ✓ 1818ms | ✓ 1748ms | http |
| 222.102.86.137:3040 | ✓ 1812ms | ✓ 1386ms | ✓ 1598ms | 否 | 否 | http |
| 101.32.243.189:80 | ✓ 1377ms | 否 | ✓ 1658ms | ✓ 1542ms | 否 | http |
| 152.42.177.32:8888 | ✓ 1259ms | 否 | ✓ 1591ms | ✓ 1530ms | 否 | http |
| 112.209.63.126:8082 | ✓ 1711ms | 否 | ✓ 1771ms | ✓ 1785ms | 否 | http |
| 117.236.124.166:3128 | ✓ 1030ms | 否 | ✓ 1012ms | ✓ 1980ms | 否 | http |
| 43.133.44.89:8888 | ✓ 1002ms | 否 | ✓ 1858ms | ✓ 1226ms | ✓ 1173ms | http |
| 38.188.247.12:999 | ✓ 1821ms | 否 | ✓ 1291ms | 否 | ✓ 1640ms | http |
| 3.101.133.120:80 | ✓ 1021ms | ✓ 1496ms | ✓ 1804ms | ✓ 1022ms | ✓ 1098ms | http |
| 118.113.244.186:1080 | ✓ 1491ms | ✓ 1858ms | ✓ 1544ms | ✓ 1757ms | ✓ 1685ms | http |
| 103.35.190.69:1082 | ✓ 180ms | ✓ 918ms | ✓ 1336ms | ✓ 949ms | ✓ 753ms | http |
| 121.230.8.136:1080 | ✓ 1338ms | ✓ 1460ms | ✓ 1079ms | ✓ 1602ms | ✓ 1600ms | http |
| 150.249.255.91:3128 | ✓ 986ms | 否 | ✓ 734ms | 否 | ✓ 829ms | http |
| 89.208.106.138:10808 | 否 | ✓ 1467ms | ✓ 467ms | 否 | ✓ 1166ms | http |
| 178.156.224.42:3128 | 否 | 否 | ✓ 1724ms | ✓ 1969ms | ✓ 1951ms | http |
| 150.107.140.238:3128 | ✓ 1015ms | 否 | ✓ 1215ms | 否 | ✓ 1406ms | http |
| 8.211.166.184:8081 | ✓ 642ms | ✓ 1296ms | ✓ 841ms | ✓ 1041ms | ✓ 896ms | http |
| 204.157.251.178:999 | ✓ 1666ms | 否 | ✓ 584ms | ✓ 1343ms | 否 | http |
| 217.182.195.221:30003 | ✓ 1829ms | 否 | ✓ 1368ms | 否 | ✓ 1921ms | http |
| 116.63.160.98:8899 | ✓ 1191ms | ✓ 1361ms | ✓ 1905ms | ✓ 1956ms | 否 | http |
| 121.230.8.250:1080 | ✓ 1248ms | ✓ 1554ms | ✓ 1516ms | ✓ 1632ms | ✓ 1382ms | http |
| 61.52.131.172:8443 | ✓ 1046ms | ✓ 1409ms | ✓ 1191ms | ✓ 1438ms | ✓ 1141ms | http |

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
