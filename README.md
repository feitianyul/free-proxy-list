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

最后更新：2026-03-14 15:31:10 UTC（2026-03-14 23:31:10 UTC+8）

**代理总数：46**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 46 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 46 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 38.145.218.82:8443 | ✓ 534ms | ✓ 1914ms | ✓ 1326ms | ✓ 1214ms | ✓ 759ms | http |
| 14.225.212.37:7890 | ✓ 1004ms | ✓ 1588ms | ✓ 1027ms | ✓ 1955ms | ✓ 980ms | http |
| 38.55.107.254:6005 | ✓ 1360ms | 否 | ✓ 1156ms | ✓ 1400ms | ✓ 845ms | http |
| 113.160.132.26:8080 | ✓ 1590ms | 否 | ✓ 1771ms | ✓ 1388ms | ✓ 1305ms | http |
| 109.73.195.10:8888 | ✓ 1661ms | ✓ 1935ms | ✓ 1511ms | 否 | 否 | http |
| 45.167.124.52:8080 | ✓ 947ms | 否 | ✓ 1556ms | 否 | ✓ 1447ms | http |
| 47.101.149.27:9010 | ✓ 1498ms | ✓ 1749ms | 否 | ✓ 1913ms | 否 | http |
| 205.209.118.30:3138 | ✓ 1018ms | 否 | ✓ 1655ms | 否 | ✓ 1892ms | http |
| 164.90.151.28:3128 | ✓ 492ms | ✓ 1362ms | ✓ 1314ms | ✓ 888ms | ✓ 684ms | http |
| 210.223.44.230:3128 | ✓ 796ms | ✓ 1424ms | ✓ 849ms | ✓ 1058ms | ✓ 811ms | http |
| 165.227.5.10:8888 | 否 | ✓ 1879ms | ✓ 924ms | 否 | ✓ 915ms | http |
| 150.230.249.50:1080 | ✓ 1058ms | 否 | 否 | ✓ 1182ms | ✓ 917ms | http |
| 35.225.22.61:80 | ✓ 260ms | ✓ 1601ms | 否 | 否 | ✓ 973ms | http |
| 62.60.177.204:34094 | ✓ 294ms | ✓ 1327ms | 否 | ✓ 956ms | ✓ 744ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1680ms | 否 | ✓ 1675ms | ✓ 1076ms | http |
| 49.51.244.112:8888 | ✓ 1331ms | ✓ 912ms | ✓ 699ms | ✓ 852ms | ✓ 647ms | http |
| 45.136.131.39:8443 | 否 | 否 | ✓ 875ms | ✓ 939ms | ✓ 739ms | http |
| 34.101.184.164:3128 | ✓ 1675ms | 否 | ✓ 1336ms | ✓ 1471ms | ✓ 1208ms | http |
| 101.43.255.96:80 | ✓ 1784ms | 否 | ✓ 1475ms | ✓ 1363ms | 否 | http |
| 45.136.131.42:8447 | ✓ 558ms | 否 | ✓ 322ms | ✓ 1137ms | ✓ 749ms | http |
| 85.198.96.242:3128 | ✓ 1639ms | 否 | ✓ 1969ms | ✓ 1674ms | ✓ 1294ms | http |
| 38.145.203.135:8443 | ✓ 822ms | ✓ 1727ms | ✓ 1926ms | ✓ 982ms | ✓ 921ms | http |
| 81.70.169.194:80 | ✓ 1395ms | 否 | ✓ 1650ms | 否 | ✓ 1473ms | http |
| 138.124.53.25:7443 | ✓ 1496ms | 否 | ✓ 1217ms | 否 | ✓ 1724ms | http |
| 62.113.119.14:8080 | ✓ 818ms | 否 | ✓ 1099ms | ✓ 1542ms | ✓ 1077ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1408ms | 否 | ✓ 1413ms | ✓ 1333ms | http |
| 103.235.67.190:80 | ✓ 1697ms | 否 | ✓ 1173ms | ✓ 1438ms | 否 | http |
| 38.145.208.186:8443 | ✓ 650ms | 否 | ✓ 535ms | ✓ 1443ms | ✓ 1520ms | http |
| 139.162.44.152:3128 | ✓ 1263ms | 否 | ✓ 892ms | ✓ 1185ms | ✓ 1147ms | http |
| 101.43.127.100:8877 | 否 | ✓ 1270ms | ✓ 1416ms | 否 | ✓ 1045ms | http |
| 101.47.73.135:3128 | ✓ 1700ms | 否 | 否 | ✓ 1769ms | ✓ 1340ms | http |
| 168.235.110.63:3128 | ✓ 656ms | ✓ 1642ms | ✓ 918ms | ✓ 1767ms | ✓ 861ms | http |
| 86.53.183.16:1080 | ✓ 993ms | 否 | ✓ 995ms | 否 | ✓ 1470ms | http |
| 91.107.148.58:53967 | ✓ 997ms | ✓ 1998ms | ✓ 1013ms | 否 | 否 | http |
| 103.84.95.54:7890 | ✓ 928ms | 否 | ✓ 1142ms | ✓ 1244ms | 否 | http |
| 113.176.92.71:3128 | ✓ 1876ms | ✓ 1713ms | ✓ 1436ms | ✓ 1657ms | ✓ 1054ms | http |
| 45.136.198.40:3128 | 否 | 否 | ✓ 1986ms | ✓ 1809ms | ✓ 1783ms | http |
| 106.14.203.63:3333 | 否 | ✓ 1211ms | ✓ 1605ms | ✓ 1244ms | ✓ 1906ms | http |
| 38.145.208.204:8443 | 否 | ✓ 1167ms | ✓ 507ms | ✓ 1091ms | ✓ 735ms | http |
| 137.184.6.37:3128 | ✓ 1134ms | ✓ 845ms | ✓ 1112ms | ✓ 943ms | ✓ 682ms | http |
| 38.55.106.245:6005 | 否 | 否 | ✓ 872ms | ✓ 1381ms | ✓ 831ms | http |
| 38.55.106.206:6005 | ✓ 1244ms | 否 | ✓ 1208ms | 否 | ✓ 1375ms | http |
| 162.240.154.26:3128 | 否 | 否 | ✓ 1961ms | ✓ 1987ms | ✓ 1918ms | http |
| 106.117.208.101:7890 | ✓ 1380ms | ✓ 1557ms | ✓ 1606ms | 否 | 否 | http |
| 47.77.193.180:1080 | 否 | ✓ 892ms | ✓ 513ms | ✓ 833ms | ✓ 626ms | http |
| 172.212.68.37:3128 | ✓ 664ms | ✓ 1417ms | ✓ 1115ms | ✓ 1238ms | ✓ 1385ms | http |

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
