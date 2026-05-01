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

最后更新：2026-05-01 15:50:03 UTC（2026-05-01 23:50:03 UTC+8）

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
| 8.219.97.248:80 | ✓ 1669ms | 否 | ✓ 1445ms | ✓ 1989ms | 否 | http |
| 206.206.126.177:2412 | ✓ 935ms | 否 | ✓ 1636ms | ✓ 1205ms | ✓ 972ms | http |
| 1.231.81.166:3128 | ✓ 1237ms | 否 | ✓ 1861ms | ✓ 1601ms | ✓ 1074ms | http |
| 113.160.132.26:8080 | ✓ 1760ms | 否 | ✓ 1510ms | ✓ 1460ms | ✓ 1144ms | http |
| 45.167.124.71:999 | ✓ 1502ms | ✓ 1680ms | ✓ 1327ms | ✓ 1702ms | ✓ 1456ms | http |
| 80.92.204.47:1081 | ✓ 1463ms | ✓ 1155ms | ✓ 1086ms | ✓ 1967ms | ✓ 1240ms | http |
| 45.140.147.82:1081 | ✓ 1791ms | ✓ 1549ms | ✓ 1203ms | 否 | ✓ 1424ms | http |
| 20.78.118.91:8561 | ✓ 1697ms | ✓ 1472ms | ✓ 1045ms | ✓ 1106ms | ✓ 934ms | http |
| 43.133.44.89:8888 | 否 | 否 | ✓ 1045ms | ✓ 1214ms | ✓ 948ms | http |
| 20.78.26.206:8561 | ✓ 1714ms | 否 | ✓ 811ms | ✓ 1132ms | ✓ 876ms | http |
| 20.210.39.153:8561 | ✓ 1729ms | 否 | ✓ 777ms | ✓ 1166ms | ✓ 882ms | http |
| 103.157.200.126:3128 | ✓ 1834ms | 否 | ✓ 1877ms | ✓ 1589ms | ✓ 1264ms | http |
| 159.223.225.118:8888 | ✓ 1308ms | 否 | ✓ 1372ms | ✓ 1776ms | ✓ 1362ms | http |
| 45.153.231.229:8080 | ✓ 718ms | 否 | ✓ 1057ms | ✓ 1633ms | ✓ 1317ms | http |
| 46.101.95.183:8888 | 否 | 否 | ✓ 919ms | ✓ 1850ms | ✓ 1521ms | http |
| 72.11.151.159:6005 | ✓ 1111ms | ✓ 1198ms | ✓ 1462ms | 否 | 否 | http |
| 47.85.51.197:1080 | 否 | 否 | ✓ 1466ms | ✓ 1009ms | ✓ 711ms | http |
| 20.127.128.70:8080 | ✓ 948ms | 否 | ✓ 782ms | ✓ 1457ms | ✓ 1295ms | http |
| 34.96.238.40:8080 | 否 | 否 | ✓ 1270ms | ✓ 1236ms | ✓ 1289ms | http |
| 2.27.40.180:1080 | ✓ 1562ms | 否 | ✓ 588ms | ✓ 1569ms | 否 | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1330ms | ✓ 1679ms | ✓ 1526ms | http |
| 86.104.72.219:1081 | ✓ 280ms | ✓ 1655ms | ✓ 79ms | 否 | ✓ 1687ms | http |
| 62.113.119.14:8080 | ✓ 605ms | ✓ 1398ms | ✓ 672ms | ✓ 1415ms | ✓ 1066ms | http |
| 86.104.72.220:1081 | ✓ 1523ms | ✓ 1936ms | 否 | ✓ 1565ms | ✓ 1778ms | http |
| 150.107.140.238:3128 | ✓ 1434ms | 否 | ✓ 1299ms | ✓ 1519ms | ✓ 1427ms | http |
| 152.32.132.190:7890 | 否 | 否 | ✓ 901ms | ✓ 1434ms | ✓ 1843ms | http |
| 43.167.237.94:3128 | ✓ 1889ms | ✓ 1296ms | ✓ 1441ms | ✓ 1014ms | ✓ 1385ms | http |
| 223.84.151.86:30005 | ✓ 1520ms | ✓ 1544ms | ✓ 1493ms | 否 | 否 | http |
| 148.230.4.241:999 | ✓ 729ms | ✓ 1659ms | ✓ 640ms | ✓ 1913ms | ✓ 1766ms | http |
| 86.104.72.220:1082 | ✓ 1375ms | ✓ 936ms | ✓ 649ms | ✓ 1058ms | ✓ 797ms | http |
| 8.154.21.175:3128 | ✓ 1731ms | ✓ 1327ms | ✓ 1056ms | ✓ 1302ms | ✓ 1117ms | http |
| 92.119.127.211:6005 | ✓ 566ms | ✓ 1915ms | 否 | ✓ 1681ms | ✓ 1969ms | http |
| 20.27.13.35:8561 | 否 | 否 | ✓ 1409ms | ✓ 1643ms | ✓ 1205ms | http |
| 20.27.11.248:8561 | 否 | ✓ 1945ms | ✓ 1263ms | ✓ 1849ms | ✓ 1201ms | http |
| 20.27.15.111:8561 | 否 | 否 | ✓ 1444ms | ✓ 1635ms | ✓ 1210ms | http |
| 20.27.14.220:8561 | ✓ 671ms | ✓ 1631ms | ✓ 796ms | ✓ 1126ms | ✓ 955ms | http |
| 20.164.75.153:8080 | ✓ 1524ms | 否 | ✓ 1914ms | 否 | ✓ 1814ms | http |
| 183.238.3.150:7897 | 否 | 否 | ✓ 1116ms | ✓ 1335ms | ✓ 1023ms | http |
| 183.232.248.73:7890 | ✓ 1084ms | ✓ 1368ms | ✓ 1158ms | ✓ 1304ms | ✓ 1028ms | http |
| 3.101.133.120:80 | 否 | ✓ 1435ms | ✓ 1098ms | 否 | ✓ 1234ms | http |
| 72.11.150.178:6005 | ✓ 1042ms | 否 | ✓ 1236ms | ✓ 1256ms | ✓ 1184ms | http |
| 37.187.109.70:10111 | ✓ 1254ms | 否 | ✓ 1992ms | 否 | ✓ 1821ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1646ms | ✓ 1270ms | 否 | ✓ 1306ms | http |
| 8.209.238.110:47701 | 否 | 否 | ✓ 856ms | ✓ 1276ms | ✓ 1495ms | http |
| 45.140.147.155:1082 | ✓ 1309ms | ✓ 1448ms | ✓ 1961ms | 否 | 否 | http |
| 154.64.232.35:8080 | ✓ 729ms | ✓ 806ms | ✓ 941ms | ✓ 1653ms | ✓ 659ms | http |
| 91.184.241.12:443 | ✓ 1647ms | ✓ 1864ms | ✓ 1155ms | ✓ 1574ms | 否 | http |
| 103.82.23.118:5196 | 否 | 否 | ✓ 1674ms | ✓ 1954ms | ✓ 1745ms | http |
| 210.223.44.230:3128 | ✓ 1680ms | 否 | ✓ 1162ms | 否 | ✓ 1934ms | http |
| 86.104.74.110:1081 | ✓ 1503ms | 否 | ✓ 1502ms | ✓ 1388ms | 否 | http |
| 121.230.8.136:1080 | ✓ 1541ms | ✓ 1816ms | 否 | ✓ 1612ms | ✓ 1272ms | http |

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
