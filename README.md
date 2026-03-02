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

最后更新：2026-03-02 15:44:04 UTC（2026-03-02 23:44:04 UTC+8）

**代理总数：55**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 55 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 55 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 103.84.95.54:7890 | ✓ 625ms | 否 | ✓ 1209ms | ✓ 825ms | ✓ 614ms | http |
| 205.209.118.30:3138 | ✓ 946ms | 否 | ✓ 1078ms | ✓ 1369ms | 否 | http |
| 5.101.0.233:3128 | ✓ 946ms | 否 | ✓ 1699ms | 否 | ✓ 1754ms | http |
| 35.225.22.61:80 | ✓ 995ms | ✓ 1285ms | ✓ 766ms | ✓ 1099ms | 否 | http |
| 47.101.149.27:9010 | 否 | 否 | ✓ 1438ms | ✓ 1405ms | ✓ 1240ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1163ms | ✓ 1722ms | 否 | ✓ 951ms | http |
| 61.72.110.54:3128 | ✓ 1507ms | 否 | ✓ 1858ms | 否 | ✓ 1514ms | http |
| 186.148.180.46:999 | ✓ 1323ms | 否 | ✓ 1462ms | ✓ 1841ms | ✓ 1947ms | http |
| 120.92.212.16:8890 | ✓ 1588ms | 否 | 否 | ✓ 1379ms | ✓ 1183ms | http |
| 142.171.85.32:1080 | ✓ 526ms | 否 | ✓ 923ms | ✓ 747ms | 否 | http |
| 114.214.162.73:26001 | ✓ 1357ms | ✓ 1588ms | ✓ 1460ms | 否 | 否 | http |
| 115.231.181.40:8128 | ✓ 1674ms | ✓ 1068ms | ✓ 911ms | 否 | 否 | http |
| 195.123.209.48:3128 | ✓ 1220ms | ✓ 1739ms | 否 | 否 | ✓ 1987ms | http |
| 81.70.169.194:80 | ✓ 977ms | ✓ 1206ms | ✓ 1615ms | ✓ 1369ms | ✓ 1270ms | http |
| 35.234.17.221:8080 | ✓ 754ms | ✓ 1084ms | 否 | ✓ 919ms | ✓ 767ms | http |
| 115.76.5.32:10010 | ✓ 1833ms | 否 | 否 | ✓ 1692ms | ✓ 1419ms | http |
| 137.184.14.135:3128 | ✓ 255ms | ✓ 1035ms | ✓ 221ms | ✓ 711ms | ✓ 483ms | http |
| 61.72.221.194:3128 | ✓ 1922ms | 否 | ✓ 1410ms | 否 | ✓ 1431ms | http |
| 101.43.255.96:80 | ✓ 905ms | 否 | ✓ 971ms | ✓ 1802ms | 否 | http |
| 103.236.64.247:8888 | ✓ 1181ms | 否 | 否 | ✓ 1470ms | ✓ 934ms | http |
| 42.203.37.113:10808 | ✓ 1264ms | 否 | ✓ 1062ms | 否 | ✓ 998ms | http |
| 210.223.44.230:3128 | ✓ 1040ms | 否 | ✓ 719ms | ✓ 876ms | ✓ 656ms | http |
| 125.128.12.14:3128 | ✓ 1043ms | ✓ 1813ms | 否 | 否 | ✓ 1115ms | http |
| 37.27.100.102:443 | ✓ 1408ms | 否 | 否 | ✓ 1910ms | ✓ 1542ms | http |
| 115.76.5.32:10009 | 否 | 否 | ✓ 1182ms | ✓ 1498ms | ✓ 1288ms | http |
| 45.125.67.37:8443 | ✓ 968ms | 否 | ✓ 852ms | ✓ 1076ms | ✓ 816ms | http |
| 115.76.5.32:10006 | 否 | 否 | ✓ 1400ms | ✓ 1598ms | ✓ 1421ms | http |
| 45.136.198.40:3128 | ✓ 1132ms | ✓ 1952ms | 否 | ✓ 1850ms | ✓ 1756ms | http |
| 162.240.154.26:3128 | ✓ 868ms | 否 | 否 | ✓ 1649ms | ✓ 1312ms | http |
| 115.76.5.32:10008 | ✓ 1686ms | 否 | ✓ 1128ms | ✓ 1507ms | ✓ 1295ms | http |
| 121.128.121.54:3128 | 否 | ✓ 1778ms | ✓ 1121ms | ✓ 1033ms | ✓ 817ms | http |
| 121.230.8.61:1080 | ✓ 1274ms | ✓ 1331ms | ✓ 959ms | 否 | 否 | http |
| 14.56.107.244:3128 | ✓ 624ms | 否 | ✓ 976ms | ✓ 974ms | ✓ 923ms | http |
| 61.72.221.234:3128 | ✓ 1896ms | 否 | ✓ 1850ms | ✓ 1026ms | ✓ 800ms | http |
| 125.128.12.194:3128 | ✓ 678ms | 否 | ✓ 1522ms | 否 | ✓ 1260ms | http |
| 125.128.12.114:3128 | ✓ 1150ms | ✓ 1590ms | ✓ 939ms | 否 | ✓ 813ms | http |
| 115.76.5.32:10005 | 否 | 否 | ✓ 1479ms | ✓ 1694ms | ✓ 1481ms | http |
| 150.249.255.91:3128 | 否 | ✓ 1522ms | 否 | ✓ 963ms | ✓ 674ms | http |
| 61.72.110.94:3128 | 否 | ✓ 857ms | 否 | ✓ 1013ms | ✓ 1265ms | http |
| 74.48.78.224:2080 | 否 | ✓ 1833ms | ✓ 968ms | ✓ 1471ms | ✓ 784ms | http |
| 64.181.240.152:3128 | ✓ 1482ms | 否 | 否 | ✓ 1461ms | ✓ 1280ms | http |
| 138.124.53.25:7443 | ✓ 1786ms | ✓ 1963ms | ✓ 1261ms | ✓ 1571ms | ✓ 1487ms | http |
| 103.39.51.190:8080 | ✓ 1704ms | 否 | 否 | ✓ 1315ms | ✓ 1266ms | http |
| 103.113.70.189:1081 | 否 | ✓ 1139ms | 否 | ✓ 1375ms | ✓ 980ms | http |
| 91.238.104.171:2023 | ✓ 1235ms | 否 | ✓ 1899ms | ✓ 1817ms | 否 | http |
| 14.56.177.44:3128 | 否 | 否 | ✓ 1536ms | ✓ 1065ms | ✓ 822ms | http |
| 200.125.170.108:999 | ✓ 870ms | ✓ 1948ms | ✓ 1727ms | 否 | 否 | http |
| 113.165.202.31:1003 | 否 | 否 | ✓ 1923ms | ✓ 1631ms | ✓ 1574ms | http |
| 115.76.5.32:10007 | 否 | 否 | ✓ 1146ms | ✓ 1435ms | ✓ 1281ms | http |
| 172.212.68.37:3128 | ✓ 1503ms | 否 | ✓ 1033ms | ✓ 1682ms | ✓ 1430ms | http |
| 120.55.163.237:10086 | ✓ 1065ms | ✓ 1011ms | ✓ 919ms | ✓ 1061ms | ✓ 821ms | http |
| 220.197.44.36:3128 | 否 | ✓ 1527ms | ✓ 1676ms | 否 | ✓ 1625ms | http |
| 103.215.36.88:15968 | 否 | ✓ 1630ms | ✓ 1693ms | 否 | ✓ 1016ms | http |
| 59.46.216.131:30001 | ✓ 1123ms | ✓ 1968ms | 否 | 否 | ✓ 1096ms | http |
| 91.238.104.172:2024 | ✓ 817ms | 否 | ✓ 1250ms | ✓ 1755ms | ✓ 1451ms | http |

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
