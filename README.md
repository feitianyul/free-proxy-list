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

最后更新：2026-03-03 10:42:59 UTC（2026-03-03 18:42:59 UTC+8）

**代理总数：44**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 44 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 44 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 3.213.157.4:3128 | ✓ 113ms | 否 | ✓ 733ms | ✓ 1487ms | ✓ 1146ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 902ms | ✓ 1228ms | ✓ 885ms | http |
| 77.83.203.6:443 | ✓ 755ms | 否 | ✓ 1958ms | 否 | ✓ 1319ms | http |
| 77.83.203.5:443 | ✓ 793ms | 否 | ✓ 1957ms | 否 | ✓ 1319ms | http |
| 138.124.53.25:7443 | ✓ 745ms | ✓ 1856ms | ✓ 1886ms | 否 | 否 | http |
| 186.148.180.46:999 | ✓ 751ms | ✓ 1843ms | ✓ 1690ms | 否 | 否 | http |
| 35.234.17.221:8080 | 否 | ✓ 1443ms | ✓ 1258ms | ✓ 1276ms | 否 | http |
| 192.166.82.55:1080 | ✓ 1292ms | 否 | 否 | ✓ 1285ms | ✓ 1479ms | http |
| 45.95.0.142:443 | ✓ 828ms | 否 | ✓ 1374ms | 否 | ✓ 1337ms | http |
| 81.70.169.194:80 | ✓ 1390ms | ✓ 1482ms | ✓ 1199ms | ✓ 1430ms | ✓ 1254ms | http |
| 109.73.195.10:8888 | ✓ 982ms | 否 | ✓ 1948ms | 否 | ✓ 1759ms | http |
| 160.238.65.4:3128 | ✓ 1467ms | 否 | ✓ 732ms | 否 | ✓ 953ms | http |
| 190.9.109.205:999 | ✓ 1594ms | 否 | ✓ 1082ms | ✓ 1240ms | ✓ 1274ms | http |
| 190.9.109.199:999 | ✓ 1593ms | 否 | ✓ 1183ms | ✓ 1186ms | ✓ 1227ms | http |
| 185.115.74.185:8080 | ✓ 841ms | ✓ 1489ms | ✓ 1885ms | 否 | 否 | http |
| 62.113.119.14:8080 | ✓ 1009ms | 否 | ✓ 542ms | 否 | ✓ 1748ms | http |
| 101.43.255.96:80 | ✓ 1631ms | ✓ 1790ms | ✓ 1505ms | ✓ 1869ms | ✓ 1456ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 1192ms | ✓ 1165ms | ✓ 1001ms | http |
| 166.0.192.117:8888 | ✓ 492ms | 否 | 否 | ✓ 1396ms | ✓ 1778ms | http |
| 120.92.212.16:7890 | ✓ 1247ms | 否 | ✓ 1962ms | 否 | ✓ 1824ms | http |
| 61.72.110.54:3128 | ✓ 1696ms | 否 | ✓ 1421ms | 否 | ✓ 986ms | http |
| 205.209.118.30:3138 | ✓ 116ms | 否 | 否 | ✓ 1048ms | ✓ 809ms | http |
| 115.231.181.40:8128 | ✓ 1159ms | ✓ 1474ms | ✓ 1153ms | 否 | 否 | http |
| 178.156.224.42:3128 | ✓ 951ms | 否 | ✓ 1844ms | ✓ 1921ms | 否 | http |
| 142.171.85.32:1080 | 否 | 否 | ✓ 778ms | ✓ 1287ms | ✓ 1364ms | http |
| 91.193.240.157:9877 | ✓ 1181ms | ✓ 1893ms | 否 | 否 | ✓ 1742ms | http |
| 157.0.142.246:10058 | 否 | ✓ 1456ms | ✓ 1163ms | ✓ 1487ms | ✓ 1113ms | http |
| 38.180.226.51:3128 | 否 | 否 | ✓ 1262ms | ✓ 1508ms | ✓ 1194ms | http |
| 5.75.196.26:40000 | ✓ 994ms | ✓ 1623ms | ✓ 681ms | 否 | ✓ 1603ms | http |
| 112.137.170.11:9401 | ✓ 1910ms | 否 | ✓ 1464ms | 否 | ✓ 1801ms | http |
| 91.217.76.97:1080 | ✓ 1022ms | 否 | ✓ 1788ms | ✓ 1931ms | 否 | http |
| 94.158.49.82:3128 | ✓ 1521ms | 否 | ✓ 1508ms | 否 | ✓ 1793ms | http |
| 45.136.198.40:3128 | ✓ 1929ms | ✓ 1488ms | ✓ 1527ms | ✓ 1989ms | ✓ 1393ms | http |
| 202.129.206.239:3128 | ✓ 1891ms | 否 | ✓ 1791ms | ✓ 1922ms | 否 | http |
| 121.230.8.11:1080 | 否 | 否 | ✓ 1989ms | ✓ 1730ms | ✓ 1265ms | http |
| 115.76.5.32:10007 | ✓ 1830ms | 否 | 否 | ✓ 1887ms | ✓ 1997ms | http |
| 223.16.170.103:3128 | 否 | 否 | ✓ 1270ms | ✓ 1343ms | ✓ 1430ms | http |
| 103.39.51.190:8080 | ✓ 1972ms | 否 | 否 | ✓ 1670ms | ✓ 1667ms | http |
| 106.14.203.63:3333 | ✓ 1019ms | ✓ 1306ms | ✓ 1129ms | ✓ 1366ms | ✓ 1065ms | http |
| 91.233.223.147:3128 | ✓ 1497ms | ✓ 1897ms | ✓ 1107ms | ✓ 1808ms | 否 | http |
| 160.238.65.9:3128 | 否 | ✓ 1266ms | ✓ 602ms | ✓ 1216ms | 否 | http |
| 160.238.65.6:3128 | 否 | ✓ 1399ms | ✓ 503ms | ✓ 1217ms | 否 | http |
| 138.197.68.35:4857 | ✓ 1576ms | ✓ 1437ms | ✓ 1114ms | 否 | 否 | http |
| 150.107.140.238:3128 | ✓ 1805ms | 否 | ✓ 1539ms | 否 | ✓ 1113ms | http |

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
