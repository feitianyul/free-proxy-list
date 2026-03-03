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

最后更新：2026-03-03 08:40:58 UTC（2026-03-03 16:40:58 UTC+8）

**代理总数：47**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 47 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 47 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 3.213.157.4:3128 | ✓ 734ms | ✓ 1585ms | ✓ 1209ms | ✓ 1835ms | 否 | http |
| 103.84.95.54:7890 | ✓ 1042ms | 否 | ✓ 853ms | 否 | ✓ 1003ms | http |
| 162.240.154.26:3128 | ✓ 737ms | 否 | ✓ 1443ms | ✓ 1672ms | ✓ 1324ms | http |
| 121.230.8.114:1080 | 否 | ✓ 1605ms | ✓ 1509ms | 否 | ✓ 1227ms | http |
| 121.128.121.54:3128 | ✓ 895ms | 否 | ✓ 1185ms | 否 | ✓ 1867ms | http |
| 91.238.104.171:2023 | 否 | 否 | ✓ 1050ms | ✓ 1701ms | ✓ 1333ms | http |
| 166.0.192.117:8888 | ✓ 998ms | 否 | 否 | ✓ 1980ms | ✓ 1747ms | http |
| 115.231.181.40:8128 | ✓ 1705ms | 否 | ✓ 1807ms | ✓ 1152ms | ✓ 871ms | http |
| 125.128.12.14:3128 | ✓ 1569ms | 否 | ✓ 1627ms | ✓ 994ms | 否 | http |
| 103.177.199.101:1111 | 否 | 否 | ✓ 1205ms | ✓ 1345ms | ✓ 1438ms | http |
| 81.70.169.194:80 | ✓ 998ms | ✓ 1537ms | ✓ 1003ms | ✓ 1289ms | ✓ 946ms | http |
| 120.92.212.16:7890 | ✓ 1894ms | 否 | ✓ 1232ms | ✓ 1919ms | 否 | http |
| 138.124.53.25:7443 | ✓ 1136ms | 否 | ✓ 1239ms | 否 | ✓ 1733ms | http |
| 101.43.255.96:80 | 否 | ✓ 1739ms | ✓ 1318ms | ✓ 1987ms | 否 | http |
| 115.76.5.32:10010 | ✓ 1455ms | 否 | 否 | ✓ 1703ms | ✓ 1726ms | http |
| 160.238.65.7:3128 | ✓ 1426ms | 否 | 否 | ✓ 1460ms | ✓ 1175ms | http |
| 160.238.65.2:3128 | 否 | 否 | ✓ 1433ms | ✓ 1464ms | ✓ 1538ms | http |
| 103.166.185.54:3128 | ✓ 1402ms | 否 | ✓ 1087ms | ✓ 1157ms | ✓ 910ms | http |
| 61.72.110.94:3128 | ✓ 1574ms | 否 | ✓ 1346ms | 否 | ✓ 765ms | http |
| 14.56.107.244:3128 | ✓ 1618ms | 否 | ✓ 1612ms | 否 | ✓ 736ms | http |
| 160.238.65.4:3128 | ✓ 1664ms | ✓ 1763ms | 否 | 否 | ✓ 1605ms | http |
| 205.209.118.30:3138 | ✓ 1412ms | 否 | ✓ 1916ms | 否 | ✓ 1027ms | http |
| 61.72.221.194:3128 | ✓ 1230ms | ✓ 1821ms | 否 | 否 | ✓ 1164ms | http |
| 35.234.17.221:8080 | ✓ 815ms | 否 | ✓ 1306ms | ✓ 1024ms | 否 | http |
| 91.193.240.157:9877 | ✓ 1251ms | 否 | ✓ 1555ms | 否 | ✓ 1526ms | http |
| 47.94.228.56:8090 | ✓ 1432ms | ✓ 1255ms | 否 | ✓ 1634ms | 否 | http |
| 61.72.110.54:3128 | ✓ 1711ms | 否 | ✓ 1292ms | ✓ 1336ms | 否 | http |
| 35.225.22.61:80 | ✓ 1165ms | 否 | ✓ 1188ms | 否 | ✓ 986ms | http |
| 5.75.196.26:40000 | ✓ 1228ms | ✓ 1852ms | ✓ 1852ms | 否 | 否 | http |
| 172.212.68.37:3128 | ✓ 1550ms | 否 | ✓ 1003ms | ✓ 1728ms | ✓ 980ms | http |
| 8.219.97.248:80 | ✓ 1588ms | 否 | ✓ 1921ms | ✓ 1899ms | 否 | http |
| 45.136.198.40:3128 | 否 | 否 | ✓ 1685ms | ✓ 1996ms | ✓ 1737ms | http |
| 45.88.0.115:3128 | ✓ 1699ms | 否 | ✓ 1704ms | 否 | ✓ 1878ms | http |
| 103.215.36.88:15556 | ✓ 1378ms | 否 | 否 | ✓ 1494ms | ✓ 1158ms | http |
| 95.85.252.153:21064 | ✓ 611ms | 否 | ✓ 1343ms | ✓ 1750ms | 否 | http |
| 120.92.212.16:8890 | ✓ 944ms | ✓ 1265ms | ✓ 1615ms | 否 | ✓ 1694ms | http |
| 62.113.119.14:8080 | ✓ 1339ms | 否 | ✓ 1112ms | ✓ 1631ms | 否 | http |
| 34.101.184.164:3128 | ✓ 1778ms | 否 | ✓ 1796ms | ✓ 1856ms | ✓ 1372ms | http |
| 202.129.206.239:3128 | 否 | 否 | ✓ 1769ms | ✓ 1513ms | ✓ 1426ms | http |
| 154.90.48.209:9090 | 否 | 否 | ✓ 1645ms | ✓ 1248ms | ✓ 1179ms | http |
| 91.238.104.172:2024 | ✓ 1361ms | ✓ 1760ms | 否 | 否 | ✓ 1931ms | http |
| 5.101.0.233:3128 | ✓ 917ms | ✓ 1781ms | ✓ 1523ms | 否 | ✓ 1828ms | http |
| 45.88.0.111:3128 | 否 | 否 | ✓ 1679ms | ✓ 1552ms | ✓ 1621ms | http |
| 14.56.177.44:3128 | ✓ 881ms | ✓ 1020ms | ✓ 687ms | ✓ 954ms | ✓ 897ms | http |
| 103.39.51.190:8080 | ✓ 1772ms | 否 | 否 | ✓ 1396ms | ✓ 1422ms | http |
| 165.225.113.220:11462 | ✓ 1588ms | 否 | 否 | ✓ 1025ms | ✓ 830ms | http |
| 61.72.221.94:3128 | ✓ 1498ms | 否 | 否 | ✓ 1220ms | ✓ 1185ms | http |

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
