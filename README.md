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

最后更新：2026-02-27 09:41:43 UTC（2026-02-27 17:41:43 UTC+8）

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
| 205.209.118.30:3138 | ✓ 1123ms | 否 | ✓ 1206ms | ✓ 1355ms | ✓ 1097ms | http |
| 147.45.216.148:1080 | ✓ 637ms | 否 | ✓ 1479ms | 否 | ✓ 1221ms | http |
| 138.124.53.25:7443 | ✓ 679ms | 否 | ✓ 1240ms | 否 | ✓ 1552ms | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1424ms | ✓ 1479ms | ✓ 1712ms | http |
| 195.123.209.48:3128 | ✓ 1263ms | 否 | ✓ 1824ms | 否 | ✓ 1957ms | http |
| 59.46.216.131:30001 | ✓ 930ms | 否 | ✓ 1270ms | ✓ 1591ms | 否 | http |
| 45.140.147.155:1081 | 否 | 否 | ✓ 1105ms | ✓ 1646ms | ✓ 1317ms | http |
| 14.56.107.244:3128 | ✓ 1668ms | ✓ 1704ms | ✓ 1783ms | 否 | 否 | http |
| 103.84.95.54:7890 | ✓ 1412ms | 否 | ✓ 667ms | 否 | ✓ 638ms | http |
| 34.142.0.1:10808 | ✓ 1075ms | 否 | ✓ 1390ms | 否 | ✓ 1619ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1289ms | ✓ 955ms | 否 | ✓ 1004ms | http |
| 101.43.255.96:80 | ✓ 1136ms | ✓ 1494ms | 否 | ✓ 1461ms | ✓ 966ms | http |
| 121.237.181.137:8888 | ✓ 1325ms | ✓ 1055ms | 否 | 否 | ✓ 938ms | http |
| 81.70.169.194:80 | 否 | ✓ 1903ms | ✓ 1051ms | ✓ 1251ms | ✓ 1067ms | http |
| 185.246.90.163:10808 | ✓ 1895ms | ✓ 1611ms | 否 | 否 | ✓ 1890ms | http |
| 111.79.111.126:3128 | ✓ 1434ms | 否 | 否 | ✓ 1402ms | ✓ 1536ms | http |
| 81.177.48.54:2080 | ✓ 1456ms | 否 | ✓ 1955ms | 否 | ✓ 1969ms | http |
| 52.188.28.218:3128 | ✓ 1317ms | 否 | ✓ 271ms | ✓ 1199ms | ✓ 873ms | http |
| 54.88.116.133:80 | ✓ 946ms | 否 | ✓ 995ms | ✓ 1625ms | ✓ 1288ms | http |
| 120.46.152.136:3128 | 否 | ✓ 1381ms | ✓ 908ms | ✓ 1399ms | ✓ 1175ms | http |
| 188.166.208.168:9876 | ✓ 752ms | 否 | ✓ 810ms | ✓ 1058ms | ✓ 848ms | http |
| 103.82.23.118:5249 | 否 | 否 | ✓ 1168ms | ✓ 1949ms | ✓ 1245ms | http |
| 168.235.110.63:3128 | ✓ 1878ms | 否 | ✓ 916ms | ✓ 1600ms | ✓ 1556ms | http |
| 35.234.17.221:8080 | ✓ 849ms | ✓ 1770ms | 否 | 否 | ✓ 1071ms | http |
| 34.205.52.219:80 | ✓ 430ms | 否 | ✓ 960ms | 否 | ✓ 852ms | http |
| 62.113.119.14:8080 | ✓ 871ms | ✓ 1858ms | ✓ 736ms | 否 | ✓ 1426ms | http |
| 115.231.181.40:8128 | ✓ 975ms | 否 | ✓ 1030ms | ✓ 1199ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1009ms | ✓ 1259ms | ✓ 933ms | ✓ 1230ms | ✓ 994ms | http |
| 207.180.228.55:80 | 否 | ✓ 1687ms | ✓ 1095ms | ✓ 1687ms | ✓ 1574ms | http |
| 121.230.8.208:1080 | 否 | ✓ 1632ms | ✓ 1479ms | ✓ 1899ms | 否 | http |
| 61.72.110.94:3128 | ✓ 1337ms | 否 | ✓ 1381ms | 否 | ✓ 824ms | http |
| 61.72.110.54:3128 | ✓ 1185ms | 否 | ✓ 685ms | 否 | ✓ 892ms | http |
| 104.247.51.76:3128 | ✓ 963ms | 否 | ✓ 1393ms | ✓ 1337ms | 否 | http |
| 36.147.78.166:80 | ✓ 1756ms | 否 | ✓ 1744ms | ✓ 1958ms | 否 | http |
| 14.56.118.154:3128 | 否 | ✓ 1561ms | ✓ 1269ms | 否 | ✓ 1444ms | http |
| 35.225.22.61:80 | ✓ 1097ms | ✓ 1375ms | ✓ 555ms | ✓ 1397ms | 否 | http |
| 121.128.121.184:3128 | ✓ 1629ms | 否 | ✓ 1635ms | ✓ 1663ms | 否 | http |
| 34.101.184.164:3128 | ✓ 1876ms | 否 | ✓ 860ms | ✓ 1260ms | ✓ 1032ms | http |
| 61.72.110.24:3128 | ✓ 1667ms | ✓ 1920ms | ✓ 1289ms | 否 | 否 | http |
| 43.248.11.162:1080 | ✓ 1900ms | 否 | ✓ 899ms | ✓ 1302ms | ✓ 798ms | http |
| 45.136.198.40:3128 | ✓ 1711ms | ✓ 1957ms | ✓ 1957ms | 否 | ✓ 1912ms | http |
| 36.147.78.166:443 | 否 | 否 | ✓ 1690ms | ✓ 1903ms | ✓ 1664ms | http |
| 121.230.9.168:1080 | ✓ 970ms | ✓ 1340ms | 否 | 否 | ✓ 899ms | http |
| 207.254.71.62:8088 | ✓ 908ms | 否 | ✓ 1788ms | ✓ 1970ms | ✓ 1677ms | http |
| 45.140.147.155:1082 | ✓ 838ms | ✓ 1859ms | ✓ 578ms | 否 | ✓ 1192ms | http |
| 3.225.78.45:80 | ✓ 1270ms | ✓ 1615ms | ✓ 1654ms | 否 | ✓ 1491ms | http |
| 103.39.51.190:8080 | ✓ 1944ms | 否 | 否 | ✓ 1831ms | ✓ 1975ms | http |
| 8.219.97.248:80 | ✓ 1286ms | 否 | ✓ 1004ms | 否 | ✓ 1914ms | http |
| 103.104.99.89:80 | ✓ 1933ms | 否 | 否 | ✓ 1544ms | ✓ 1285ms | http |
| 103.117.100.127:13082 | ✓ 639ms | 否 | ✓ 655ms | ✓ 826ms | ✓ 646ms | http |
| 103.67.46.225:3125 | ✓ 1451ms | 否 | ✓ 1803ms | ✓ 1541ms | ✓ 1590ms | http |
| 85.208.108.43:2094 | 否 | 否 | ✓ 574ms | ✓ 1169ms | ✓ 1026ms | http |

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
