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

最后更新：2026-05-03 09:41:36 UTC（2026-05-03 17:41:36 UTC+8）

**代理总数：52**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 52 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 52 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 113.160.132.26:8080 | ✓ 1898ms | ✓ 1737ms | 否 | ✓ 1458ms | ✓ 1258ms | http |
| 45.167.124.71:999 | 否 | 否 | ✓ 1508ms | ✓ 1882ms | ✓ 1461ms | http |
| 86.104.72.219:1081 | ✓ 1073ms | ✓ 1243ms | ✓ 932ms | 否 | ✓ 1081ms | http |
| 80.92.204.47:1081 | ✓ 455ms | ✓ 1267ms | ✓ 1314ms | ✓ 1985ms | ✓ 1495ms | http |
| 206.206.126.177:2412 | ✓ 826ms | ✓ 1743ms | ✓ 1120ms | ✓ 1138ms | ✓ 953ms | http |
| 86.104.72.220:1081 | 否 | 否 | ✓ 813ms | ✓ 1518ms | ✓ 1314ms | http |
| 86.104.72.220:1082 | 否 | ✓ 1921ms | ✓ 1202ms | ✓ 1236ms | 否 | http |
| 109.120.156.122:8090 | ✓ 931ms | 否 | ✓ 585ms | 否 | ✓ 1903ms | http |
| 193.123.250.39:1080 | ✓ 1056ms | 否 | ✓ 1198ms | ✓ 1655ms | ✓ 1083ms | http |
| 168.110.52.228:3128 | ✓ 907ms | 否 | ✓ 1354ms | ✓ 1386ms | ✓ 1251ms | http |
| 38.180.62.47:10808 | ✓ 805ms | ✓ 1107ms | 否 | ✓ 1476ms | ✓ 1358ms | http |
| 37.187.109.70:10111 | ✓ 1410ms | ✓ 1975ms | ✓ 1389ms | 否 | ✓ 1826ms | http |
| 45.153.231.229:8080 | 否 | ✓ 1832ms | ✓ 1516ms | 否 | ✓ 1857ms | http |
| 154.64.232.35:8080 | ✓ 1563ms | ✓ 1522ms | 否 | ✓ 1324ms | 否 | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 953ms | ✓ 1294ms | ✓ 975ms | http |
| 120.92.108.86:7890 | ✓ 1437ms | 否 | 否 | ✓ 1895ms | ✓ 1600ms | http |
| 103.153.149.138:8080 | ✓ 1534ms | 否 | ✓ 1267ms | 否 | ✓ 1383ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1095ms | ✓ 1553ms | ✓ 1902ms | ✓ 1272ms | http |
| 47.85.51.197:1080 | ✓ 1398ms | ✓ 1280ms | ✓ 622ms | 否 | 否 | http |
| 43.133.44.89:8888 | ✓ 1048ms | 否 | ✓ 1857ms | 否 | ✓ 1602ms | http |
| 59.46.216.131:30001 | ✓ 1185ms | ✓ 1965ms | ✓ 1277ms | ✓ 1476ms | 否 | http |
| 20.164.75.153:8080 | ✓ 1772ms | 否 | ✓ 1943ms | 否 | ✓ 1928ms | http |
| 62.133.60.126:24558 | ✓ 1630ms | ✓ 1841ms | ✓ 1152ms | 否 | 否 | http |
| 8.154.21.175:3128 | ✓ 971ms | 否 | ✓ 1089ms | ✓ 1222ms | ✓ 1037ms | http |
| 45.125.67.37:8443 | ✓ 1072ms | 否 | ✓ 1287ms | ✓ 1280ms | ✓ 1093ms | http |
| 46.105.190.38:3128 | ✓ 998ms | ✓ 1585ms | ✓ 1647ms | 否 | ✓ 1587ms | http |
| 135.125.97.184:43205 | ✓ 1387ms | ✓ 1891ms | ✓ 1733ms | 否 | 否 | http |
| 121.135.144.141:8096 | ✓ 1060ms | ✓ 1243ms | ✓ 1354ms | 否 | ✓ 1832ms | http |
| 180.103.19.181:1080 | 否 | 否 | ✓ 1248ms | ✓ 1492ms | ✓ 1320ms | http |
| 86.104.74.110:1082 | ✓ 489ms | ✓ 1710ms | ✓ 520ms | ✓ 1595ms | ✓ 1448ms | http |
| 86.104.74.110:1081 | ✓ 428ms | ✓ 1906ms | ✓ 675ms | ✓ 1301ms | ✓ 1449ms | http |
| 212.58.132.5:8888 | ✓ 1162ms | 否 | ✓ 1330ms | ✓ 1492ms | ✓ 1251ms | http |
| 46.105.190.40:3128 | ✓ 972ms | ✓ 1350ms | ✓ 961ms | ✓ 1721ms | 否 | http |
| 152.42.177.32:8888 | ✓ 1280ms | 否 | ✓ 1308ms | ✓ 1457ms | ✓ 1484ms | http |
| 154.90.48.209:9090 | ✓ 1887ms | 否 | ✓ 1477ms | ✓ 1818ms | ✓ 1412ms | http |
| 120.92.212.16:7890 | ✓ 1075ms | ✓ 1335ms | ✓ 1966ms | ✓ 1302ms | ✓ 1037ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1298ms | ✓ 1099ms | ✓ 1736ms | ✓ 1767ms | http |
| 8.219.97.248:80 | ✓ 1352ms | 否 | 否 | ✓ 1472ms | ✓ 1305ms | http |
| 3.101.133.120:80 | ✓ 823ms | ✓ 1295ms | ✓ 869ms | ✓ 1353ms | ✓ 1226ms | http |
| 94.131.118.129:1082 | ✓ 432ms | ✓ 1184ms | 否 | ✓ 1481ms | ✓ 1550ms | http |
| 91.217.81.131:1080 | ✓ 1094ms | 否 | ✓ 1176ms | 否 | ✓ 1566ms | http |
| 94.131.118.39:1081 | ✓ 1015ms | ✓ 1467ms | ✓ 1052ms | ✓ 1778ms | ✓ 1163ms | http |
| 106.10.55.212:1121 | ✓ 1918ms | ✓ 1734ms | 否 | ✓ 1641ms | 否 | http |
| 34.96.238.40:8080 | ✓ 1339ms | ✓ 1223ms | 否 | 否 | ✓ 1166ms | http |
| 103.82.23.118:5207 | ✓ 1652ms | 否 | ✓ 1558ms | ✓ 1625ms | ✓ 1651ms | http |
| 176.100.39.18:8080 | ✓ 1362ms | ✓ 1824ms | ✓ 1878ms | 否 | ✓ 1248ms | http |
| 155.212.131.230:3128 | ✓ 1890ms | 否 | ✓ 1998ms | 否 | ✓ 1880ms | http |
| 94.131.118.129:1081 | ✓ 1250ms | 否 | ✓ 683ms | ✓ 1585ms | 否 | http |
| 160.238.65.3:3128 | ✓ 1067ms | ✓ 1426ms | ✓ 898ms | ✓ 1448ms | ✓ 1043ms | http |
| 160.238.65.9:3128 | ✓ 1067ms | ✓ 1517ms | ✓ 808ms | ✓ 1448ms | ✓ 1040ms | http |
| 160.238.65.4:3128 | ✓ 1067ms | 否 | ✓ 538ms | ✓ 1520ms | ✓ 998ms | http |
| 160.238.65.7:3128 | ✓ 1068ms | 否 | ✓ 536ms | ✓ 1524ms | ✓ 998ms | http |

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
