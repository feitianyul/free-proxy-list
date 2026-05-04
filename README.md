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

最后更新：2026-05-04 12:45:30 UTC（2026-05-04 20:45:30 UTC+8）

**代理总数：49**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 49 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 49 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 34.71.229.255:3128 | ✓ 596ms | ✓ 1895ms | ✓ 798ms | ✓ 1168ms | ✓ 950ms | http |
| 181.119.97.24:999 | ✓ 1534ms | ✓ 1792ms | ✓ 1964ms | 否 | ✓ 1599ms | http |
| 206.206.126.177:2412 | ✓ 948ms | 否 | ✓ 1467ms | ✓ 1311ms | ✓ 994ms | http |
| 137.59.47.73:3128 | ✓ 1171ms | 否 | ✓ 1135ms | ✓ 1277ms | ✓ 1018ms | http |
| 45.140.147.82:1081 | ✓ 1528ms | 否 | ✓ 1549ms | ✓ 1952ms | ✓ 1895ms | http |
| 45.140.147.82:1082 | ✓ 1489ms | 否 | ✓ 1539ms | ✓ 1966ms | ✓ 1881ms | http |
| 113.160.132.26:8080 | ✓ 1189ms | ✓ 1669ms | 否 | ✓ 1532ms | ✓ 1156ms | http |
| 34.101.184.164:3128 | ✓ 1772ms | 否 | ✓ 1146ms | ✓ 1986ms | ✓ 1168ms | http |
| 47.85.51.197:1080 | 否 | ✓ 878ms | ✓ 911ms | 否 | ✓ 696ms | http |
| 91.233.223.147:3128 | ✓ 1417ms | 否 | ✓ 1367ms | 否 | ✓ 1517ms | http |
| 103.157.200.126:3128 | ✓ 1175ms | 否 | ✓ 1795ms | 否 | ✓ 1882ms | http |
| 46.105.190.40:3128 | ✓ 1158ms | 否 | ✓ 505ms | 否 | ✓ 1294ms | http |
| 120.92.212.16:7890 | ✓ 1173ms | ✓ 1405ms | ✓ 1159ms | 否 | ✓ 1497ms | http |
| 171.234.50.242:5116 | ✓ 1514ms | 否 | ✓ 1720ms | ✓ 1850ms | ✓ 1625ms | http |
| 120.92.212.16:8890 | ✓ 1672ms | ✓ 1839ms | 否 | 否 | ✓ 1596ms | http |
| 86.104.72.220:1081 | ✓ 214ms | ✓ 891ms | ✓ 390ms | ✓ 1373ms | 否 | http |
| 38.188.247.12:999 | 否 | 否 | ✓ 539ms | ✓ 1704ms | ✓ 1634ms | http |
| 45.59.122.132:80 | ✓ 1205ms | 否 | ✓ 879ms | 否 | ✓ 1183ms | http |
| 45.153.231.229:8080 | ✓ 824ms | 否 | ✓ 1317ms | 否 | ✓ 1723ms | http |
| 101.6.65.112:10080 | 否 | ✓ 1591ms | ✓ 1261ms | ✓ 1736ms | ✓ 1244ms | http |
| 193.123.250.39:1080 | ✓ 1515ms | 否 | 否 | ✓ 1561ms | ✓ 1588ms | http |
| 38.180.62.47:10808 | ✓ 866ms | ✓ 1771ms | 否 | ✓ 1055ms | 否 | http |
| 59.46.216.131:30001 | ✓ 1354ms | ✓ 1582ms | ✓ 1382ms | 否 | ✓ 1399ms | http |
| 62.133.60.126:24558 | ✓ 1032ms | ✓ 1618ms | 否 | 否 | ✓ 1890ms | http |
| 80.92.204.47:1082 | 否 | ✓ 1329ms | ✓ 1420ms | ✓ 1334ms | ✓ 1500ms | http |
| 8.154.21.175:3128 | ✓ 980ms | ✓ 1317ms | ✓ 1024ms | ✓ 1349ms | ✓ 1088ms | http |
| 160.250.134.143:3128 | ✓ 1634ms | 否 | ✓ 1168ms | ✓ 1453ms | ✓ 1233ms | http |
| 218.108.131.186:17890 | ✓ 1467ms | ✓ 1651ms | 否 | 否 | ✓ 1351ms | http |
| 31.56.48.253:26133 | ✓ 1484ms | 否 | ✓ 890ms | ✓ 1994ms | ✓ 1522ms | http |
| 154.64.232.35:8080 | ✓ 392ms | 否 | ✓ 1198ms | ✓ 1319ms | ✓ 900ms | http |
| 107.173.42.121:7890 | 否 | ✓ 1389ms | ✓ 710ms | ✓ 1707ms | 否 | http |
| 3.101.133.120:80 | ✓ 613ms | ✓ 1834ms | ✓ 1544ms | ✓ 1191ms | ✓ 1235ms | http |
| 86.104.74.110:1081 | ✓ 1037ms | ✓ 1982ms | ✓ 857ms | ✓ 1537ms | ✓ 1033ms | http |
| 45.173.12.140:1994 | ✓ 1043ms | 否 | ✓ 838ms | ✓ 1582ms | ✓ 1361ms | http |
| 190.12.150.244:999 | ✓ 940ms | ✓ 1638ms | ✓ 934ms | 否 | 否 | http |
| 38.180.2.107:3128 | ✓ 1372ms | ✓ 1912ms | ✓ 1803ms | 否 | 否 | http |
| 144.91.102.48:3128 | ✓ 859ms | ✓ 1358ms | ✓ 455ms | ✓ 1874ms | ✓ 1108ms | http |
| 109.120.156.122:8090 | ✓ 1525ms | 否 | ✓ 870ms | ✓ 1698ms | 否 | http |
| 77.110.107.80:8080 | ✓ 697ms | ✓ 1846ms | 否 | 否 | ✓ 1470ms | http |
| 152.70.91.193:40000 | 否 | 否 | ✓ 1836ms | ✓ 1490ms | ✓ 1271ms | http |
| 152.32.132.190:7890 | ✓ 1425ms | ✓ 1762ms | ✓ 1178ms | ✓ 1676ms | ✓ 1354ms | http |
| 45.186.6.104:3128 | ✓ 1784ms | ✓ 1752ms | ✓ 1745ms | 否 | 否 | http |
| 86.104.74.110:1082 | ✓ 1242ms | 否 | ✓ 432ms | ✓ 1134ms | ✓ 1116ms | http |
| 80.92.204.47:1081 | ✓ 1163ms | 否 | ✓ 493ms | 否 | ✓ 1463ms | http |
| 207.254.71.62:8088 | ✓ 694ms | ✓ 1633ms | ✓ 1114ms | ✓ 1853ms | ✓ 1765ms | http |
| 120.92.108.86:7890 | ✓ 1716ms | 否 | 否 | ✓ 1886ms | ✓ 1567ms | http |
| 77.110.107.80:1080 | ✓ 893ms | 否 | ✓ 1738ms | ✓ 1605ms | 否 | http |
| 61.52.131.172:8443 | ✓ 1168ms | ✓ 1453ms | ✓ 1166ms | ✓ 1445ms | ✓ 1167ms | http |
| 103.172.70.173:8080 | ✓ 1612ms | 否 | 否 | ✓ 1768ms | ✓ 1749ms | http |

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
