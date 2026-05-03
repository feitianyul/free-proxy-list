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

最后更新：2026-05-03 16:42:03 UTC（2026-05-04 00:42:03 UTC+8）

**代理总数：58**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 58 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 58 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 20.78.118.91:8561 | ✓ 919ms | ✓ 1407ms | ✓ 901ms | ✓ 1463ms | ✓ 1095ms | http |
| 20.210.39.153:8561 | ✓ 922ms | ✓ 1151ms | ✓ 1012ms | ✓ 1328ms | ✓ 1371ms | http |
| 154.64.232.35:8080 | 否 | ✓ 1124ms | 否 | ✓ 1301ms | ✓ 810ms | http |
| 20.78.26.206:8561 | ✓ 911ms | 否 | ✓ 910ms | ✓ 1313ms | ✓ 1229ms | http |
| 206.206.126.177:2412 | ✓ 1184ms | 否 | ✓ 1500ms | ✓ 1255ms | ✓ 950ms | http |
| 1.231.81.166:3128 | ✓ 1050ms | ✓ 1871ms | 否 | ✓ 1297ms | ✓ 1441ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1869ms | ✓ 1735ms | ✓ 1826ms | 否 | http |
| 152.32.132.190:7890 | 否 | 否 | ✓ 913ms | ✓ 1143ms | ✓ 905ms | http |
| 34.101.184.164:3128 | ✓ 1980ms | 否 | ✓ 1814ms | 否 | ✓ 1793ms | http |
| 148.230.4.241:999 | ✓ 878ms | 否 | ✓ 802ms | ✓ 1633ms | ✓ 1335ms | http |
| 45.167.124.71:999 | ✓ 1179ms | 否 | ✓ 1308ms | ✓ 1665ms | ✓ 1867ms | http |
| 94.131.118.39:1081 | 否 | 否 | ✓ 756ms | ✓ 1906ms | ✓ 1125ms | http |
| 193.23.209.44:3128 | ✓ 459ms | 否 | ✓ 948ms | ✓ 1513ms | ✓ 1293ms | http |
| 212.58.132.5:8888 | ✓ 1130ms | 否 | ✓ 1815ms | ✓ 1519ms | ✓ 1172ms | http |
| 91.233.223.147:3128 | ✓ 1918ms | ✓ 1862ms | 否 | 否 | ✓ 1946ms | http |
| 94.131.118.129:1081 | ✓ 976ms | ✓ 1396ms | ✓ 363ms | ✓ 1490ms | ✓ 1065ms | http |
| 45.78.79.225:1080 | ✓ 991ms | 否 | ✓ 1433ms | 否 | ✓ 1592ms | http |
| 46.105.190.38:3128 | ✓ 985ms | 否 | ✓ 858ms | ✓ 1689ms | 否 | http |
| 47.85.51.197:1080 | ✓ 712ms | 否 | ✓ 1591ms | ✓ 922ms | ✓ 1662ms | http |
| 117.236.124.166:3128 | ✓ 1586ms | 否 | ✓ 1581ms | 否 | ✓ 1852ms | http |
| 62.133.60.126:24558 | 否 | 否 | ✓ 503ms | ✓ 1409ms | ✓ 1224ms | http |
| 46.105.190.40:3128 | ✓ 815ms | 否 | ✓ 1087ms | 否 | ✓ 1787ms | http |
| 86.104.72.219:1081 | ✓ 361ms | 否 | ✓ 1477ms | 否 | ✓ 1270ms | http |
| 193.123.250.39:1080 | ✓ 926ms | 否 | ✓ 1459ms | ✓ 1559ms | ✓ 1265ms | http |
| 8.219.97.248:80 | ✓ 1963ms | 否 | ✓ 1433ms | ✓ 1837ms | ✓ 1692ms | http |
| 130.61.174.200:1080 | ✓ 464ms | 否 | ✓ 1529ms | ✓ 1913ms | 否 | http |
| 121.130.199.80:3128 | ✓ 989ms | ✓ 1487ms | ✓ 1797ms | ✓ 1740ms | ✓ 1468ms | http |
| 38.180.62.47:10808 | 否 | ✓ 1798ms | ✓ 1545ms | 否 | ✓ 805ms | http |
| 190.12.150.244:999 | ✓ 1046ms | 否 | ✓ 1243ms | ✓ 1710ms | ✓ 1455ms | http |
| 43.133.44.89:8888 | ✓ 934ms | 否 | ✓ 879ms | 否 | ✓ 1806ms | http |
| 62.113.119.14:8080 | 否 | ✓ 1578ms | 否 | ✓ 1357ms | ✓ 1105ms | http |
| 86.104.72.220:1082 | ✓ 1562ms | ✓ 915ms | ✓ 294ms | ✓ 973ms | ✓ 1779ms | http |
| 45.125.67.37:8443 | ✓ 1077ms | 否 | ✓ 1076ms | ✓ 1311ms | ✓ 1454ms | http |
| 8.154.21.175:3128 | ✓ 1998ms | 否 | ✓ 917ms | ✓ 1097ms | ✓ 950ms | http |
| 34.96.238.40:8080 | 否 | ✓ 1567ms | ✓ 1258ms | 否 | ✓ 1318ms | http |
| 103.82.23.118:5236 | ✓ 1793ms | 否 | ✓ 1528ms | ✓ 1795ms | ✓ 1910ms | http |
| 220.197.44.36:3128 | ✓ 1467ms | 否 | ✓ 1296ms | ✓ 1990ms | ✓ 1685ms | http |
| 109.120.156.122:8090 | ✓ 1562ms | 否 | ✓ 848ms | 否 | ✓ 1670ms | http |
| 5.161.196.81:8888 | ✓ 1326ms | 否 | ✓ 977ms | ✓ 886ms | ✓ 745ms | http |
| 144.31.25.69:21064 | ✓ 1069ms | 否 | ✓ 892ms | 否 | ✓ 1267ms | http |
| 218.108.131.186:17890 | ✓ 811ms | ✓ 1049ms | ✓ 854ms | ✓ 1088ms | ✓ 923ms | http |
| 20.164.75.153:8080 | ✓ 1558ms | 否 | ✓ 1673ms | 否 | ✓ 1529ms | http |
| 80.92.204.47:1081 | ✓ 428ms | 否 | ✓ 1264ms | ✓ 1710ms | ✓ 1283ms | http |
| 45.140.147.82:1081 | ✓ 587ms | ✓ 1376ms | ✓ 1495ms | 否 | ✓ 1371ms | http |
| 120.92.108.86:7890 | ✓ 1638ms | 否 | ✓ 1731ms | 否 | ✓ 1428ms | http |
| 59.46.216.131:30001 | ✓ 1055ms | 否 | ✓ 1282ms | 否 | ✓ 1138ms | http |
| 86.104.72.220:1081 | ✓ 1716ms | ✓ 871ms | ✓ 49ms | ✓ 1170ms | 否 | http |
| 64.188.67.154:1080 | ✓ 1299ms | ✓ 1593ms | ✓ 1025ms | ✓ 1630ms | 否 | http |
| 178.156.224.42:3128 | ✓ 954ms | 否 | ✓ 1794ms | ✓ 1897ms | 否 | http |
| 104.128.138.186:1080 | ✓ 1741ms | 否 | ✓ 1778ms | ✓ 1958ms | ✓ 1788ms | http |
| 150.107.140.238:3128 | ✓ 1915ms | 否 | 否 | ✓ 1404ms | ✓ 1159ms | http |
| 154.90.48.209:9090 | ✓ 1750ms | 否 | ✓ 1102ms | ✓ 1502ms | ✓ 1152ms | http |
| 101.32.243.189:80 | ✓ 1731ms | 否 | ✓ 1566ms | ✓ 1753ms | ✓ 1506ms | http |
| 121.230.9.198:1080 | ✓ 1152ms | ✓ 1956ms | ✓ 1172ms | 否 | 否 | http |
| 125.76.214.178:8091 | ✓ 1459ms | ✓ 1716ms | ✓ 1208ms | 否 | ✓ 1164ms | http |
| 103.217.216.65:8181 | ✓ 1656ms | 否 | ✓ 1801ms | ✓ 1680ms | 否 | http |
| 61.52.131.172:8443 | ✓ 875ms | ✓ 1193ms | ✓ 1016ms | 否 | ✓ 1616ms | http |
| 3.101.133.120:80 | ✓ 1510ms | ✓ 1362ms | ✓ 1660ms | 否 | ✓ 1360ms | http |

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
