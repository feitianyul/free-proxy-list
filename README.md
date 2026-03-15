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

最后更新：2026-03-15 11:27:08 UTC（2026-03-15 19:27:08 UTC+8）

**代理总数：40**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 39 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 40 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 38.145.218.82:8443 | ✓ 663ms | ✓ 1550ms | ✓ 1103ms | ✓ 1105ms | ✓ 700ms | http |
| 113.160.132.26:8080 | 否 | 否 | ✓ 1385ms | ✓ 1383ms | ✓ 1113ms | http |
| 115.231.181.40:8128 | 否 | 否 | ✓ 1154ms | ✓ 1334ms | ✓ 1488ms | http |
| 45.167.124.52:8080 | ✓ 1741ms | 否 | ✓ 1425ms | 否 | ✓ 1574ms | http |
| 205.209.118.30:3138 | 否 | ✓ 949ms | ✓ 753ms | ✓ 970ms | ✓ 724ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1408ms | ✓ 1132ms | ✓ 1425ms | ✓ 1154ms | http |
| 178.236.245.17:3128 | ✓ 978ms | 否 | ✓ 977ms | ✓ 1929ms | 否 | http |
| 95.3.9.78:8080 | ✓ 1111ms | ✓ 1737ms | ✓ 1300ms | ✓ 1586ms | ✓ 1226ms | http |
| 81.70.169.194:80 | 否 | ✓ 1933ms | ✓ 1244ms | ✓ 1458ms | ✓ 1625ms | http |
| 8.219.97.248:80 | ✓ 1529ms | 否 | ✓ 1289ms | ✓ 1677ms | ✓ 1582ms | http |
| 120.92.212.16:7890 | ✓ 1118ms | ✓ 1424ms | 否 | ✓ 1789ms | ✓ 1129ms | http |
| 35.225.22.61:80 | 否 | ✓ 1316ms | 否 | ✓ 1135ms | ✓ 979ms | http |
| 137.220.150.152:6005 | ✓ 1696ms | 否 | ✓ 1066ms | ✓ 1358ms | ✓ 1461ms | http |
| 165.227.5.10:8888 | ✓ 298ms | ✓ 1778ms | 否 | 否 | ✓ 951ms | http |
| 95.3.9.78:3128 | ✓ 1485ms | 否 | 否 | ✓ 1553ms | ✓ 1192ms | http |
| 101.43.255.96:80 | ✓ 1128ms | ✓ 1485ms | ✓ 1617ms | ✓ 1413ms | ✓ 1173ms | http |
| 137.220.151.110:6005 | ✓ 932ms | 否 | ✓ 1076ms | ✓ 1302ms | ✓ 1033ms | http |
| 212.192.13.76:6005 | ✓ 1792ms | 否 | 否 | ✓ 1958ms | ✓ 1264ms | http |
| 137.220.150.104:6005 | ✓ 1037ms | 否 | ✓ 1148ms | ✓ 1264ms | ✓ 1151ms | http |
| 45.136.130.223:8443 | ✓ 382ms | ✓ 1457ms | ✓ 525ms | ✓ 1698ms | ✓ 773ms | http |
| 38.145.203.135:8443 | ✓ 485ms | ✓ 927ms | ✓ 915ms | ✓ 913ms | ✓ 909ms | http |
| 45.136.131.28:8447 | 否 | 否 | ✓ 961ms | ✓ 942ms | ✓ 817ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 868ms | ✓ 1034ms | ✓ 835ms | http |
| 88.80.150.82:8080 | ✓ 1269ms | 否 | ✓ 1055ms | ✓ 1905ms | 否 | https |
| 139.162.46.62:3128 | ✓ 1309ms | 否 | ✓ 920ms | ✓ 1828ms | ✓ 1429ms | http |
| 38.55.106.208:6005 | ✓ 1713ms | 否 | 否 | ✓ 1211ms | ✓ 1593ms | http |
| 164.90.155.209:3128 | ✓ 1184ms | ✓ 1050ms | ✓ 1258ms | ✓ 918ms | ✓ 718ms | http |
| 101.43.127.100:8877 | ✓ 1328ms | 否 | 否 | ✓ 1936ms | ✓ 1061ms | http |
| 123.57.0.163:8888 | ✓ 1683ms | ✓ 1703ms | 否 | ✓ 1967ms | ✓ 1621ms | http |
| 2.56.122.146:10808 | ✓ 970ms | 否 | ✓ 1421ms | ✓ 1787ms | ✓ 1467ms | http |
| 45.119.85.216:3128 | ✓ 1509ms | 否 | ✓ 1670ms | ✓ 1959ms | ✓ 1164ms | http |
| 85.208.108.43:2094 | ✓ 226ms | 否 | ✓ 350ms | ✓ 917ms | ✓ 974ms | http |
| 45.136.198.40:3128 | ✓ 667ms | 否 | ✓ 1800ms | 否 | ✓ 1557ms | http |
| 113.176.92.71:3128 | 否 | ✓ 1755ms | ✓ 1834ms | ✓ 1738ms | ✓ 1511ms | http |
| 85.198.96.242:3128 | ✓ 1763ms | 否 | ✓ 922ms | 否 | ✓ 1717ms | http |
| 43.225.185.4:8000 | ✓ 1661ms | 否 | ✓ 1753ms | ✓ 1556ms | ✓ 1243ms | http |
| 103.39.51.190:8080 | ✓ 1991ms | 否 | 否 | ✓ 1876ms | ✓ 1642ms | http |
| 121.40.231.103:7890 | ✓ 1090ms | ✓ 1243ms | ✓ 1046ms | ✓ 1401ms | ✓ 1060ms | http |
| 121.147.253.205:3104 | 否 | 否 | ✓ 1455ms | ✓ 1536ms | ✓ 1393ms | http |
| 106.14.203.63:3333 | ✓ 997ms | ✓ 1242ms | 否 | 否 | ✓ 1596ms | http |

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
