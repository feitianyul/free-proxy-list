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

最后更新：2026-03-09 22:36:51 UTC（2026-03-10 06:36:51 UTC+8）

**代理总数：58**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 57 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 58 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 1.231.81.166:3128 | ✓ 1354ms | ✓ 1099ms | ✓ 1034ms | ✓ 1117ms | ✓ 924ms | http |
| 61.72.110.114:3128 | ✓ 1365ms | 否 | 否 | ✓ 1417ms | ✓ 1008ms | http |
| 202.155.12.161:443 | ✓ 1989ms | 否 | ✓ 1847ms | ✓ 1443ms | ✓ 1327ms | http |
| 14.225.212.37:7890 | 否 | ✓ 1826ms | ✓ 1213ms | 否 | ✓ 1911ms | http |
| 47.77.193.180:1080 | 否 | 否 | ✓ 458ms | ✓ 973ms | ✓ 764ms | http |
| 91.107.141.42:8081 | ✓ 953ms | ✓ 1496ms | 否 | ✓ 1538ms | ✓ 1384ms | http |
| 49.147.68.237:8082 | ✓ 1971ms | 否 | ✓ 1912ms | ✓ 1529ms | ✓ 1534ms | http |
| 62.113.119.14:8080 | ✓ 1132ms | ✓ 1534ms | ✓ 1137ms | 否 | 否 | http |
| 35.225.22.61:80 | ✓ 783ms | ✓ 1977ms | ✓ 981ms | 否 | ✓ 844ms | http |
| 154.3.236.202:3128 | ✓ 303ms | ✓ 1464ms | ✓ 1021ms | ✓ 1185ms | ✓ 935ms | http |
| 210.77.29.245:7890 | ✓ 1019ms | ✓ 1603ms | ✓ 1196ms | ✓ 1342ms | ✓ 1052ms | http |
| 101.43.255.96:80 | ✓ 1157ms | ✓ 1438ms | ✓ 1241ms | ✓ 1450ms | ✓ 1209ms | http |
| 39.104.201.40:7890 | 否 | 否 | ✓ 1158ms | ✓ 1447ms | ✓ 1112ms | http |
| 190.9.109.198:999 | ✓ 857ms | ✓ 1392ms | ✓ 1162ms | ✓ 1332ms | ✓ 1205ms | http |
| 121.237.181.137:8888 | ✓ 1006ms | ✓ 1235ms | ✓ 1004ms | ✓ 1358ms | 否 | http |
| 162.240.154.26:3128 | ✓ 900ms | ✓ 1496ms | ✓ 1676ms | ✓ 1546ms | ✓ 1223ms | http |
| 193.168.173.136:443 | ✓ 584ms | 否 | ✓ 1261ms | 否 | ✓ 1991ms | http |
| 178.236.245.17:3128 | ✓ 586ms | ✓ 1452ms | ✓ 1629ms | ✓ 1994ms | ✓ 1430ms | http |
| 178.236.245.59:3128 | ✓ 587ms | ✓ 1718ms | ✓ 1360ms | 否 | ✓ 1410ms | http |
| 81.70.169.194:80 | ✓ 1114ms | ✓ 1582ms | ✓ 1142ms | ✓ 1489ms | ✓ 1206ms | http |
| 115.231.181.40:8128 | ✓ 1013ms | 否 | ✓ 1057ms | 否 | ✓ 1121ms | http |
| 101.47.73.135:3128 | ✓ 1387ms | 否 | 否 | ✓ 1539ms | ✓ 1308ms | http |
| 116.80.49.169:3172 | ✓ 1743ms | 否 | ✓ 1882ms | 否 | ✓ 1845ms | http |
| 194.213.18.200:443 | ✓ 986ms | 否 | 否 | ✓ 1434ms | ✓ 1918ms | http |
| 152.42.213.210:8080 | ✓ 1635ms | 否 | ✓ 1513ms | ✓ 1645ms | ✓ 990ms | http |
| 220.170.182.39:9293 | ✓ 1692ms | ✓ 1630ms | ✓ 1635ms | ✓ 1634ms | ✓ 1678ms | http |
| 46.183.25.8:443 | ✓ 890ms | 否 | ✓ 927ms | ✓ 1221ms | ✓ 938ms | http |
| 210.223.44.230:3128 | ✓ 760ms | ✓ 1112ms | ✓ 818ms | ✓ 1091ms | ✓ 902ms | http |
| 103.113.70.189:1081 | 否 | ✓ 1187ms | 否 | ✓ 1523ms | ✓ 785ms | http |
| 83.219.250.8:62920 | ✓ 631ms | 否 | ✓ 1156ms | 否 | ✓ 1498ms | http |
| 113.176.92.71:3128 | 否 | ✓ 1557ms | ✓ 1719ms | ✓ 1387ms | ✓ 1127ms | http |
| 61.72.110.54:3128 | ✓ 1078ms | ✓ 1141ms | ✓ 983ms | ✓ 1787ms | ✓ 975ms | http |
| 91.233.223.147:3128 | ✓ 1036ms | 否 | ✓ 1425ms | 否 | ✓ 1969ms | http |
| 88.80.150.82:8080 | ✓ 1416ms | 否 | ✓ 1884ms | 否 | ✓ 1754ms | https |
| 165.227.5.10:8888 | ✓ 647ms | 否 | 否 | ✓ 1601ms | ✓ 1090ms | http |
| 34.101.184.164:3128 | ✓ 1771ms | 否 | ✓ 1908ms | ✓ 1760ms | ✓ 1327ms | http |
| 1.225.116.115:1080 | ✓ 1891ms | ✓ 1111ms | ✓ 916ms | ✓ 1119ms | ✓ 883ms | http |
| 205.209.118.30:3138 | ✓ 1802ms | 否 | ✓ 1999ms | 否 | ✓ 1784ms | http |
| 45.136.131.47:8443 | ✓ 330ms | ✓ 1050ms | ✓ 1356ms | ✓ 966ms | ✓ 857ms | http |
| 61.72.221.94:3128 | ✓ 1858ms | ✓ 1568ms | ✓ 1333ms | ✓ 1604ms | ✓ 1771ms | http |
| 45.136.198.40:3128 | ✓ 664ms | ✓ 1761ms | ✓ 1460ms | 否 | ✓ 1515ms | http |
| 120.92.212.16:8890 | ✓ 1195ms | 否 | ✓ 1548ms | 否 | ✓ 1928ms | http |
| 162.248.165.72:1080 | ✓ 1711ms | ✓ 1999ms | ✓ 1422ms | 否 | 否 | http |
| 5.252.33.13:2025 | 否 | 否 | ✓ 1742ms | ✓ 1968ms | ✓ 1599ms | http |
| 45.129.141.143:3128 | ✓ 911ms | 否 | ✓ 1814ms | ✓ 1746ms | ✓ 1611ms | http |
| 185.191.236.162:3128 | ✓ 954ms | 否 | ✓ 1954ms | ✓ 1905ms | ✓ 1327ms | http |
| 180.130.80.196:9003 | ✓ 1431ms | ✓ 1647ms | ✓ 1538ms | 否 | 否 | http |
| 103.39.51.190:8080 | ✓ 1679ms | 否 | ✓ 1982ms | ✓ 1932ms | ✓ 1576ms | http |
| 94.176.3.43:7443 | ✓ 1298ms | 否 | ✓ 1860ms | ✓ 1921ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1931ms | ✓ 1670ms | ✓ 1110ms | 否 | ✓ 1133ms | http |
| 138.124.53.25:7443 | ✓ 901ms | 否 | ✓ 1522ms | ✓ 1951ms | 否 | http |
| 106.14.205.114:483 | ✓ 1270ms | ✓ 1261ms | ✓ 1372ms | ✓ 1323ms | ✓ 1053ms | http |
| 203.205.49.2:10002 | ✓ 1847ms | 否 | ✓ 1730ms | ✓ 1755ms | ✓ 1533ms | http |
| 103.35.188.243:3128 | ✓ 928ms | ✓ 1580ms | 否 | ✓ 1939ms | ✓ 837ms | http |
| 113.177.131.2:3128 | ✓ 1597ms | 否 | ✓ 1302ms | ✓ 1911ms | ✓ 1253ms | http |
| 5.101.0.233:3128 | ✓ 964ms | ✓ 1623ms | ✓ 967ms | ✓ 1532ms | ✓ 1183ms | http |
| 125.128.12.144:3128 | ✓ 1592ms | 否 | ✓ 1077ms | ✓ 1357ms | ✓ 1114ms | http |
| 120.55.163.237:10086 | ✓ 1327ms | ✓ 1293ms | ✓ 1105ms | ✓ 1380ms | 否 | http |

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
