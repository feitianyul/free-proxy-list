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

最后更新：2026-04-25 09:47:40 UTC（2026-04-25 17:47:40 UTC+8）

**代理总数：51**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 51 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 51 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 113.160.132.26:8080 | 否 | ✓ 1512ms | ✓ 1374ms | ✓ 1836ms | ✓ 1106ms | http |
| 1.231.81.166:3128 | 否 | 否 | ✓ 1600ms | ✓ 1280ms | ✓ 882ms | http |
| 46.101.95.183:8888 | ✓ 1962ms | 否 | 否 | ✓ 1745ms | ✓ 1943ms | http |
| 177.93.132.244:3128 | ✓ 676ms | 否 | ✓ 697ms | 否 | ✓ 1718ms | http |
| 218.108.131.186:17890 | ✓ 1006ms | ✓ 1263ms | ✓ 968ms | ✓ 1767ms | ✓ 1070ms | http |
| 168.110.52.228:3128 | ✓ 893ms | 否 | 否 | ✓ 1473ms | ✓ 948ms | http |
| 212.58.132.5:8888 | ✓ 1231ms | 否 | ✓ 1312ms | ✓ 1546ms | ✓ 1238ms | http |
| 130.61.174.200:1080 | ✓ 479ms | 否 | 否 | ✓ 1689ms | ✓ 980ms | http |
| 2.27.54.161:1080 | 否 | 否 | ✓ 1739ms | ✓ 1758ms | ✓ 1311ms | http |
| 168.144.75.9:3128 | 否 | 否 | ✓ 1679ms | ✓ 1914ms | ✓ 1642ms | http |
| 45.140.147.155:1081 | ✓ 743ms | ✓ 1916ms | ✓ 1789ms | 否 | 否 | http |
| 45.140.147.155:1082 | ✓ 1116ms | 否 | ✓ 1328ms | ✓ 1986ms | 否 | http |
| 38.7.23.118:999 | ✓ 1527ms | ✓ 1245ms | ✓ 916ms | ✓ 1180ms | ✓ 1452ms | http |
| 47.85.51.197:1080 | ✓ 89ms | ✓ 982ms | ✓ 1116ms | 否 | 否 | http |
| 45.140.147.82:1081 | ✓ 1190ms | ✓ 1483ms | 否 | ✓ 1791ms | ✓ 1077ms | http |
| 80.92.204.47:1081 | ✓ 394ms | ✓ 1202ms | ✓ 1206ms | ✓ 1596ms | ✓ 1280ms | http |
| 166.88.61.54:8000 | ✓ 1171ms | ✓ 1386ms | ✓ 1097ms | ✓ 1047ms | ✓ 1234ms | http |
| 206.206.126.177:2412 | ✓ 898ms | ✓ 1795ms | ✓ 1091ms | 否 | 否 | http |
| 110.76.145.116:8090 | ✓ 1894ms | 否 | ✓ 1657ms | ✓ 1622ms | ✓ 1637ms | http |
| 62.113.119.14:8080 | 否 | ✓ 1378ms | ✓ 1242ms | ✓ 1436ms | ✓ 1847ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1685ms | ✓ 1388ms | ✓ 1409ms | 否 | http |
| 91.99.15.45:2095 | ✓ 834ms | 否 | ✓ 1767ms | 否 | ✓ 1878ms | http |
| 120.92.212.16:7890 | ✓ 1375ms | ✓ 1396ms | ✓ 1102ms | ✓ 1602ms | 否 | http |
| 120.92.108.86:7890 | ✓ 1708ms | 否 | 否 | ✓ 1836ms | ✓ 1654ms | http |
| 94.131.122.129:1081 | ✓ 1914ms | 否 | ✓ 1572ms | 否 | ✓ 1798ms | http |
| 162.240.154.26:3128 | ✓ 1004ms | 否 | ✓ 1236ms | ✓ 1799ms | 否 | http |
| 84.47.150.125:1080 | 否 | 否 | ✓ 1146ms | ✓ 1632ms | ✓ 1598ms | http |
| 82.114.228.67:1080 | ✓ 1296ms | 否 | ✓ 1134ms | ✓ 1562ms | 否 | http |
| 185.234.66.87:1081 | ✓ 1557ms | 否 | ✓ 964ms | ✓ 1896ms | ✓ 1329ms | http |
| 34.71.229.255:3128 | ✓ 621ms | ✓ 1475ms | ✓ 819ms | ✓ 1055ms | ✓ 802ms | http |
| 91.193.240.157:9877 | ✓ 1006ms | 否 | ✓ 1571ms | 否 | ✓ 1650ms | http |
| 35.225.22.61:80 | ✓ 848ms | ✓ 1199ms | 否 | 否 | ✓ 860ms | http |
| 86.104.74.110:1081 | ✓ 1030ms | ✓ 1431ms | ✓ 991ms | ✓ 1374ms | ✓ 1102ms | http |
| 183.232.248.73:7890 | 否 | 否 | ✓ 1070ms | ✓ 1296ms | ✓ 1971ms | http |
| 45.140.147.82:1082 | ✓ 1463ms | ✓ 1841ms | ✓ 1353ms | ✓ 1504ms | ✓ 1987ms | http |
| 45.186.6.104:3128 | ✓ 1335ms | ✓ 1911ms | ✓ 1990ms | 否 | 否 | http |
| 210.77.22.138:7890 | ✓ 1031ms | ✓ 1350ms | ✓ 1105ms | ✓ 1354ms | ✓ 1053ms | http |
| 85.190.99.143:443 | ✓ 1072ms | ✓ 1724ms | ✓ 1865ms | 否 | ✓ 1580ms | http |
| 161.35.181.96:999 | ✓ 178ms | ✓ 878ms | ✓ 52ms | ✓ 890ms | ✓ 696ms | http |
| 159.223.225.118:8888 | 否 | ✓ 1550ms | 否 | ✓ 1505ms | ✓ 1338ms | http |
| 20.120.225.109:3128 | ✓ 950ms | ✓ 1287ms | ✓ 1321ms | ✓ 1343ms | ✓ 1228ms | http |
| 152.32.132.190:7890 | 否 | ✓ 1203ms | ✓ 1260ms | ✓ 1131ms | 否 | http |
| 210.223.44.230:3128 | ✓ 1429ms | ✓ 1374ms | ✓ 888ms | ✓ 1413ms | 否 | http |
| 150.107.140.238:3128 | ✓ 962ms | 否 | ✓ 1387ms | ✓ 1400ms | ✓ 1072ms | http |
| 117.236.124.166:3128 | ✓ 1793ms | 否 | ✓ 1002ms | ✓ 1816ms | ✓ 1400ms | http |
| 64.188.77.26:3128 | ✓ 946ms | 否 | ✓ 597ms | 否 | ✓ 1024ms | http |
| 45.76.207.177:40000 | ✓ 977ms | 否 | ✓ 1072ms | ✓ 1253ms | ✓ 1096ms | http |
| 103.126.238.13:8081 | 否 | 否 | ✓ 1717ms | ✓ 1882ms | ✓ 1774ms | http |
| 61.52.131.172:8443 | ✓ 1721ms | ✓ 1407ms | ✓ 1056ms | ✓ 1453ms | 否 | http |
| 103.39.51.207:8080 | ✓ 1618ms | 否 | ✓ 1957ms | ✓ 1592ms | ✓ 1872ms | http |
| 94.131.106.231:1081 | ✓ 711ms | ✓ 1258ms | ✓ 1538ms | ✓ 1646ms | ✓ 1164ms | http |

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
