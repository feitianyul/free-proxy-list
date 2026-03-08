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

最后更新：2026-03-08 11:23:23 UTC（2026-03-08 19:23:23 UTC+8）

**代理总数：32**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 32 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 32 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | 否 | 否 | ✓ 1202ms | ✓ 1261ms | ✓ 974ms | http |
| 202.155.12.161:443 | ✓ 1876ms | 否 | 否 | ✓ 1849ms | ✓ 1392ms | http |
| 1.231.81.166:3128 | ✓ 1823ms | 否 | ✓ 1787ms | 否 | ✓ 1626ms | http |
| 120.92.212.16:8890 | ✓ 968ms | ✓ 1242ms | 否 | 否 | ✓ 997ms | http |
| 14.225.222.164:7890 | 否 | 否 | ✓ 1173ms | ✓ 1647ms | ✓ 1813ms | http |
| 120.92.212.16:7890 | ✓ 1193ms | ✓ 1205ms | 否 | 否 | ✓ 1234ms | http |
| 194.213.18.200:443 | ✓ 1882ms | 否 | 否 | ✓ 1770ms | ✓ 1221ms | http |
| 46.183.25.8:443 | ✓ 1859ms | 否 | ✓ 1193ms | ✓ 1425ms | 否 | http |
| 101.43.255.96:80 | ✓ 962ms | ✓ 1242ms | ✓ 1039ms | ✓ 1346ms | ✓ 975ms | http |
| 106.14.203.63:3333 | 否 | ✓ 1672ms | 否 | ✓ 1124ms | ✓ 1793ms | http |
| 112.78.187.186:8080 | ✓ 1952ms | 否 | ✓ 1252ms | ✓ 1437ms | ✓ 1401ms | http |
| 81.70.169.194:80 | ✓ 1279ms | ✓ 1359ms | 否 | ✓ 1248ms | 否 | http |
| 115.231.181.40:8128 | ✓ 900ms | ✓ 1782ms | ✓ 1980ms | 否 | 否 | http |
| 5.252.33.13:2025 | ✓ 1526ms | 否 | ✓ 1377ms | 否 | ✓ 1896ms | http |
| 59.46.216.131:30001 | ✓ 952ms | ✓ 1406ms | ✓ 1432ms | 否 | 否 | http |
| 35.225.22.61:80 | 否 | ✓ 1463ms | ✓ 967ms | ✓ 988ms | ✓ 1141ms | http |
| 103.215.36.88:17715 | 否 | 否 | ✓ 1228ms | ✓ 1518ms | ✓ 1972ms | http |
| 45.205.28.107:8080 | ✓ 409ms | 否 | ✓ 1286ms | ✓ 886ms | 否 | http |
| 185.243.218.43:49153 | ✓ 711ms | 否 | ✓ 1559ms | 否 | ✓ 1765ms | http |
| 62.113.119.14:8080 | ✓ 1313ms | 否 | ✓ 1107ms | ✓ 1565ms | ✓ 1249ms | http |
| 39.104.201.40:7890 | 否 | 否 | ✓ 1003ms | ✓ 1264ms | ✓ 953ms | http |
| 47.77.193.180:1080 | 否 | ✓ 1416ms | ✓ 485ms | ✓ 723ms | ✓ 568ms | http |
| 103.215.36.88:16090 | ✓ 1201ms | ✓ 1435ms | ✓ 1151ms | ✓ 1990ms | ✓ 1290ms | http |
| 162.248.165.72:1080 | ✓ 1268ms | 否 | 否 | ✓ 1879ms | ✓ 1453ms | http |
| 45.140.147.82:1081 | ✓ 1001ms | 否 | ✓ 1549ms | ✓ 1720ms | ✓ 1352ms | http |
| 45.136.198.40:3128 | ✓ 820ms | 否 | 否 | ✓ 1935ms | ✓ 1753ms | http |
| 159.89.31.62:8080 | ✓ 947ms | 否 | ✓ 1678ms | 否 | ✓ 1121ms | http |
| 45.186.6.104:3128 | ✓ 1530ms | ✓ 1686ms | ✓ 1745ms | 否 | 否 | http |
| 103.69.84.106:3131 | ✓ 1862ms | 否 | ✓ 1348ms | ✓ 1151ms | ✓ 941ms | http |
| 103.39.51.190:8080 | 否 | 否 | ✓ 1883ms | ✓ 1773ms | ✓ 1484ms | http |
| 103.67.46.225:3125 | 否 | 否 | ✓ 1568ms | ✓ 1635ms | ✓ 1548ms | http |
| 211.171.114.154:3128 | 否 | ✓ 1448ms | ✓ 1640ms | ✓ 1930ms | ✓ 1239ms | http |

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
