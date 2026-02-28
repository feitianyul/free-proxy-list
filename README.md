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

最后更新：2026-02-28 16:22:41 UTC（2026-03-01 00:22:41 UTC+8）

**代理总数：48**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 48 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 48 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 163ms | 否 | ✓ 1162ms | 否 | ✓ 1020ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1381ms | 否 | ✓ 1429ms | ✓ 1116ms | http |
| 222.228.171.92:8080 | ✓ 1889ms | 否 | ✓ 1917ms | ✓ 1558ms | ✓ 1604ms | http |
| 121.237.181.137:8888 | ✓ 1822ms | ✓ 1195ms | ✓ 1990ms | ✓ 1313ms | 否 | http |
| 59.46.216.131:30001 | 否 | ✓ 1505ms | ✓ 1114ms | ✓ 1621ms | ✓ 1159ms | http |
| 120.92.212.16:8890 | ✓ 1337ms | ✓ 1364ms | ✓ 1113ms | 否 | ✓ 1120ms | http |
| 212.175.29.184:8080 | ✓ 1926ms | 否 | ✓ 1909ms | 否 | ✓ 1864ms | http |
| 36.147.78.166:443 | 否 | 否 | ✓ 1918ms | ✓ 1767ms | ✓ 1842ms | http |
| 101.43.255.96:80 | ✓ 1486ms | 否 | 否 | ✓ 1311ms | ✓ 1100ms | http |
| 81.70.169.194:80 | ✓ 1114ms | 否 | ✓ 1821ms | ✓ 1837ms | 否 | http |
| 115.231.181.40:8128 | ✓ 1995ms | ✓ 1277ms | ✓ 1479ms | 否 | ✓ 1259ms | http |
| 91.238.104.171:2023 | ✓ 1437ms | 否 | ✓ 1759ms | 否 | ✓ 1617ms | http |
| 113.176.92.71:3128 | 否 | 否 | ✓ 1676ms | ✓ 1323ms | ✓ 1037ms | http |
| 165.225.120.17:10458 | 否 | 否 | ✓ 1772ms | ✓ 1790ms | ✓ 1367ms | http |
| 142.171.85.32:1080 | ✓ 1254ms | 否 | ✓ 1536ms | ✓ 1428ms | ✓ 1119ms | http |
| 35.225.22.61:80 | ✓ 1047ms | 否 | ✓ 298ms | ✓ 920ms | ✓ 769ms | http |
| 193.124.225.175:1080 | ✓ 1982ms | 否 | ✓ 1780ms | ✓ 1777ms | 否 | http |
| 85.208.108.43:2094 | 否 | 否 | ✓ 1475ms | ✓ 1095ms | ✓ 948ms | http |
| 148.135.85.87:1080 | ✓ 1210ms | ✓ 1834ms | 否 | ✓ 941ms | ✓ 811ms | http |
| 36.147.78.166:80 | ✓ 1848ms | 否 | 否 | ✓ 1793ms | ✓ 1764ms | http |
| 165.225.120.17:10906 | ✓ 1299ms | 否 | ✓ 904ms | ✓ 1773ms | ✓ 1431ms | http |
| 165.225.120.17:11995 | 否 | 否 | ✓ 853ms | ✓ 1792ms | ✓ 1416ms | http |
| 165.225.120.17:10880 | 否 | 否 | ✓ 851ms | ✓ 1788ms | ✓ 1644ms | http |
| 121.40.231.103:7890 | ✓ 1526ms | ✓ 1185ms | ✓ 1926ms | 否 | 否 | http |
| 3.213.157.4:3128 | ✓ 856ms | ✓ 948ms | ✓ 1872ms | ✓ 1144ms | ✓ 744ms | http |
| 85.208.108.43:10808 | ✓ 1824ms | 否 | 否 | ✓ 1322ms | ✓ 927ms | http |
| 165.225.120.17:11070 | ✓ 1745ms | ✓ 1981ms | ✓ 1276ms | 否 | 否 | http |
| 165.225.120.17:11178 | ✓ 1754ms | ✓ 1964ms | ✓ 1288ms | ✓ 1920ms | ✓ 1376ms | http |
| 165.225.120.17:12215 | ✓ 1752ms | 否 | ✓ 1251ms | 否 | ✓ 1431ms | http |
| 165.225.120.17:11099 | ✓ 1751ms | 否 | ✓ 1253ms | 否 | ✓ 1472ms | http |
| 165.225.120.17:12497 | ✓ 1754ms | 否 | ✓ 1377ms | 否 | ✓ 1424ms | http |
| 165.225.120.17:12265 | ✓ 1749ms | 否 | ✓ 1638ms | 否 | ✓ 1398ms | http |
| 31.59.129.75:8080 | ✓ 1053ms | ✓ 1670ms | 否 | 否 | ✓ 1472ms | http |
| 52.188.28.218:3128 | ✓ 142ms | 否 | ✓ 875ms | 否 | ✓ 921ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 827ms | ✓ 1045ms | ✓ 1445ms | http |
| 165.225.120.17:10919 | ✓ 1904ms | 否 | ✓ 867ms | ✓ 1700ms | ✓ 1376ms | http |
| 165.225.120.17:11912 | ✓ 1911ms | 否 | ✓ 870ms | ✓ 1720ms | ✓ 1397ms | http |
| 165.225.120.17:11702 | ✓ 846ms | 否 | ✓ 853ms | ✓ 1807ms | 否 | http |
| 165.225.120.17:11745 | ✓ 1053ms | 否 | ✓ 1273ms | ✓ 1834ms | 否 | http |
| 172.212.68.37:3128 | ✓ 850ms | ✓ 1335ms | ✓ 1978ms | 否 | 否 | http |
| 84.247.149.172:3128 | ✓ 837ms | 否 | ✓ 1446ms | ✓ 1171ms | ✓ 1643ms | http |
| 35.241.222.101:3128 | 否 | ✓ 1750ms | ✓ 1007ms | ✓ 1842ms | ✓ 1244ms | http |
| 45.136.198.40:3128 | ✓ 704ms | 否 | ✓ 990ms | ✓ 1479ms | ✓ 1207ms | http |
| 49.7.179.70:3333 | 否 | 否 | ✓ 1700ms | ✓ 1988ms | ✓ 1571ms | http |
| 121.230.9.205:1080 | ✓ 1280ms | ✓ 1928ms | 否 | ✓ 1766ms | 否 | http |
| 121.230.8.99:1080 | ✓ 1907ms | ✓ 1950ms | 否 | ✓ 1961ms | ✓ 1447ms | http |
| 103.133.223.21:8080 | 否 | 否 | ✓ 1511ms | ✓ 1587ms | ✓ 1649ms | http |
| 45.140.147.155:1081 | ✓ 437ms | 否 | ✓ 1221ms | ✓ 1841ms | ✓ 1118ms | http |

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
