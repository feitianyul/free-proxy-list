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

最后更新：2026-04-26 21:35:45 UTC（2026-04-27 05:35:45 UTC+8）

**代理总数：34**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 34 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 34 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 46.101.95.183:8888 | ✓ 1458ms | 否 | ✓ 1513ms | 否 | ✓ 1467ms | http |
| 47.84.59.16:1080 | ✓ 1565ms | ✓ 1764ms | ✓ 1052ms | ✓ 1190ms | ✓ 944ms | http |
| 120.92.212.16:7890 | ✓ 1018ms | ✓ 1518ms | ✓ 1569ms | 否 | ✓ 1889ms | http |
| 212.58.132.5:8888 | ✓ 1709ms | 否 | ✓ 1711ms | ✓ 1567ms | ✓ 1220ms | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 1308ms | ✓ 1291ms | ✓ 1490ms | http |
| 177.93.132.244:3128 | ✓ 857ms | 否 | ✓ 1880ms | 否 | ✓ 1925ms | http |
| 47.85.51.197:1080 | ✓ 88ms | 否 | ✓ 216ms | ✓ 1013ms | 否 | http |
| 113.160.132.26:8080 | 否 | ✓ 1475ms | ✓ 1156ms | ✓ 1685ms | ✓ 1146ms | http |
| 8.137.112.117:3128 | ✓ 1034ms | ✓ 1487ms | ✓ 1292ms | ✓ 1421ms | ✓ 1144ms | http |
| 152.32.132.190:7890 | 否 | 否 | ✓ 1260ms | ✓ 1038ms | ✓ 1338ms | http |
| 47.105.98.23:3128 | ✓ 972ms | 否 | ✓ 1714ms | ✓ 1591ms | 否 | http |
| 36.141.21.200:7890 | ✓ 1139ms | ✓ 1291ms | 否 | ✓ 1705ms | 否 | http |
| 103.157.200.126:3128 | ✓ 1985ms | 否 | ✓ 1827ms | ✓ 1646ms | ✓ 1344ms | http |
| 8.219.97.248:80 | ✓ 1281ms | 否 | ✓ 1498ms | ✓ 1684ms | 否 | http |
| 121.230.8.237:1080 | ✓ 1224ms | ✓ 1695ms | ✓ 1155ms | 否 | 否 | http |
| 64.188.67.154:1080 | ✓ 497ms | ✓ 1883ms | ✓ 1620ms | ✓ 1968ms | ✓ 1420ms | http |
| 121.230.8.158:1080 | ✓ 1257ms | ✓ 1792ms | ✓ 1188ms | ✓ 1481ms | ✓ 1295ms | http |
| 62.113.119.14:8080 | ✓ 628ms | 否 | ✓ 577ms | ✓ 1483ms | 否 | http |
| 1.234.153.14:80 | ✓ 1393ms | ✓ 1065ms | ✓ 1049ms | ✓ 1009ms | ✓ 1162ms | http |
| 86.104.72.219:1081 | ✓ 1361ms | ✓ 914ms | ✓ 696ms | 否 | ✓ 854ms | http |
| 47.112.25.109:7890 | ✓ 943ms | ✓ 1853ms | 否 | ✓ 1306ms | 否 | http |
| 120.92.108.86:7890 | ✓ 1270ms | 否 | ✓ 1543ms | ✓ 1728ms | ✓ 1394ms | http |
| 168.144.75.9:3128 | ✓ 1183ms | 否 | ✓ 1363ms | ✓ 1976ms | 否 | http |
| 218.108.131.186:17890 | ✓ 925ms | ✓ 1167ms | ✓ 949ms | ✓ 1216ms | ✓ 985ms | http |
| 183.232.248.73:7890 | ✓ 1343ms | 否 | ✓ 1733ms | 否 | ✓ 1619ms | http |
| 60.249.94.208:3128 | ✓ 1345ms | 否 | ✓ 873ms | 否 | ✓ 887ms | http |
| 34.101.184.164:3128 | ✓ 1817ms | 否 | ✓ 1273ms | 否 | ✓ 1227ms | http |
| 77.110.119.136:3128 | ✓ 962ms | 否 | 否 | ✓ 1114ms | ✓ 854ms | http |
| 8.211.166.184:8081 | ✓ 1307ms | ✓ 1253ms | ✓ 838ms | ✓ 991ms | ✓ 793ms | http |
| 103.147.134.133:3125 | ✓ 1524ms | 否 | 否 | ✓ 1670ms | ✓ 1646ms | http |
| 61.52.131.172:8443 | ✓ 1070ms | ✓ 1268ms | ✓ 1049ms | ✓ 1292ms | 否 | http |
| 103.162.54.251:8083 | 否 | 否 | ✓ 1567ms | ✓ 1543ms | ✓ 1510ms | http |
| 217.77.102.18:3128 | ✓ 1125ms | 否 | ✓ 1575ms | ✓ 1970ms | ✓ 1600ms | http |
| 103.39.51.207:8080 | ✓ 1553ms | 否 | ✓ 1999ms | 否 | ✓ 1515ms | http |

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
