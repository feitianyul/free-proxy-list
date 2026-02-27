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

最后更新：2026-02-27 16:38:49 UTC（2026-02-28 00:38:49 UTC+8）

**代理总数：42**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 42 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 42 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 478ms | 否 | ✓ 324ms | ✓ 1384ms | ✓ 1280ms | http |
| 35.225.22.61:80 | ✓ 722ms | 否 | ✓ 1614ms | ✓ 1149ms | ✓ 922ms | http |
| 132.145.93.138:1080 | ✓ 856ms | 否 | ✓ 1327ms | ✓ 1557ms | ✓ 1064ms | http |
| 147.45.216.148:1080 | ✓ 706ms | 否 | ✓ 1532ms | ✓ 1980ms | ✓ 1504ms | http |
| 210.223.44.230:3128 | ✓ 1570ms | ✓ 1866ms | ✓ 1875ms | ✓ 1905ms | ✓ 653ms | http |
| 52.188.28.218:3128 | ✓ 1340ms | 否 | ✓ 1335ms | 否 | ✓ 875ms | http |
| 121.237.181.137:8888 | 否 | ✓ 1709ms | ✓ 1110ms | ✓ 1096ms | ✓ 1738ms | http |
| 34.101.184.164:3128 | ✓ 1727ms | 否 | ✓ 1402ms | ✓ 1551ms | ✓ 1367ms | http |
| 14.56.107.244:3128 | ✓ 619ms | 否 | ✓ 1156ms | ✓ 1790ms | ✓ 739ms | http |
| 120.92.212.16:7890 | ✓ 1022ms | ✓ 1374ms | ✓ 1203ms | 否 | 否 | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 937ms | ✓ 1008ms | ✓ 1354ms | http |
| 85.208.108.43:2094 | ✓ 509ms | 否 | 否 | ✓ 1460ms | ✓ 1062ms | http |
| 61.72.110.24:3128 | ✓ 1165ms | 否 | 否 | ✓ 1277ms | ✓ 810ms | http |
| 120.92.212.16:8890 | ✓ 1160ms | 否 | 否 | ✓ 1228ms | ✓ 947ms | http |
| 36.147.78.166:80 | 否 | ✓ 1523ms | ✓ 1621ms | ✓ 1884ms | 否 | http |
| 36.147.78.166:443 | ✓ 1611ms | ✓ 1574ms | ✓ 1621ms | ✓ 1726ms | 否 | http |
| 168.235.110.63:3128 | ✓ 641ms | 否 | ✓ 1276ms | 否 | ✓ 1073ms | http |
| 101.47.73.135:3128 | ✓ 834ms | 否 | ✓ 1036ms | ✓ 1994ms | ✓ 826ms | http |
| 81.70.169.194:80 | ✓ 870ms | 否 | ✓ 952ms | ✓ 1181ms | ✓ 1845ms | http |
| 45.125.67.37:8443 | ✓ 1049ms | 否 | 否 | ✓ 979ms | ✓ 848ms | http |
| 101.43.255.96:80 | ✓ 918ms | 否 | 否 | ✓ 1784ms | ✓ 1001ms | http |
| 61.72.110.94:3128 | 否 | ✓ 1451ms | ✓ 1818ms | ✓ 1380ms | 否 | http |
| 207.248.113.12:8080 | ✓ 854ms | 否 | ✓ 1556ms | ✓ 1474ms | ✓ 1230ms | http |
| 120.46.152.136:3128 | ✓ 1083ms | ✓ 1125ms | ✓ 835ms | ✓ 1126ms | ✓ 864ms | http |
| 104.37.184.214:1080 | ✓ 1212ms | 否 | 否 | ✓ 787ms | ✓ 578ms | http |
| 103.113.70.189:1081 | 否 | ✓ 1425ms | 否 | ✓ 1479ms | ✓ 1570ms | http |
| 103.135.102.161:8081 | ✓ 1893ms | 否 | ✓ 1533ms | ✓ 1342ms | ✓ 1548ms | http |
| 115.231.181.40:8128 | ✓ 852ms | ✓ 1181ms | 否 | ✓ 1920ms | 否 | http |
| 138.124.53.25:7443 | ✓ 1422ms | 否 | ✓ 1649ms | ✓ 1865ms | 否 | http |
| 54.173.56.172:80 | ✓ 456ms | ✓ 1706ms | ✓ 1441ms | 否 | 否 | http |
| 59.46.216.131:30001 | ✓ 1108ms | 否 | ✓ 1819ms | ✓ 1300ms | 否 | http |
| 103.216.104.6:1080 | ✓ 1681ms | 否 | ✓ 1520ms | 否 | ✓ 1726ms | http |
| 45.140.147.82:1081 | ✓ 1186ms | 否 | ✓ 1338ms | 否 | ✓ 1167ms | http |
| 211.171.114.154:3128 | ✓ 1420ms | ✓ 1403ms | ✓ 1210ms | ✓ 1338ms | ✓ 857ms | http |
| 202.129.206.239:3128 | ✓ 1352ms | 否 | ✓ 1384ms | ✓ 1491ms | ✓ 1350ms | http |
| 8.219.97.248:80 | ✓ 1267ms | 否 | ✓ 1559ms | ✓ 1497ms | 否 | http |
| 103.82.23.118:5242 | ✓ 1568ms | 否 | ✓ 1156ms | 否 | ✓ 1290ms | http |
| 45.136.198.40:3128 | ✓ 850ms | 否 | ✓ 964ms | ✓ 1676ms | ✓ 1296ms | http |
| 98.85.203.202:80 | ✓ 417ms | 否 | ✓ 1748ms | ✓ 1699ms | 否 | http |
| 103.39.51.190:8080 | ✓ 1747ms | 否 | ✓ 1685ms | ✓ 1461ms | ✓ 1591ms | http |
| 138.197.68.35:4857 | ✓ 738ms | ✓ 1805ms | ✓ 1655ms | ✓ 1460ms | 否 | http |
| 83.219.250.8:62920 | ✓ 1180ms | 否 | ✓ 1749ms | 否 | ✓ 1930ms | http |

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
