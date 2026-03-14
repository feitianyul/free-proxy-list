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

最后更新：2026-03-14 17:28:09 UTC（2026-03-15 01:28:09 UTC+8）

**代理总数：41**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 40 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 41 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 38.145.218.82:8443 | ✓ 1191ms | ✓ 1735ms | ✓ 1022ms | ✓ 982ms | ✓ 717ms | http |
| 109.73.195.10:8888 | ✓ 1209ms | 否 | ✓ 1752ms | ✓ 1832ms | 否 | http |
| 150.230.249.50:1080 | 否 | 否 | ✓ 1223ms | ✓ 1190ms | ✓ 1970ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1666ms | ✓ 1705ms | ✓ 1464ms | 否 | http |
| 45.167.124.52:8080 | ✓ 1898ms | ✓ 1726ms | 否 | ✓ 1603ms | ✓ 1639ms | http |
| 192.71.213.85:9812 | ✓ 867ms | 否 | ✓ 1361ms | ✓ 1863ms | 否 | http |
| 35.225.22.61:80 | ✓ 523ms | 否 | ✓ 989ms | 否 | ✓ 1032ms | http |
| 81.70.169.194:80 | ✓ 1169ms | 否 | ✓ 1215ms | ✓ 1534ms | ✓ 1610ms | http |
| 101.43.255.96:80 | ✓ 1579ms | 否 | ✓ 1105ms | ✓ 1961ms | 否 | http |
| 165.232.188.102:9090 | ✓ 1872ms | 否 | ✓ 1322ms | ✓ 1524ms | ✓ 1206ms | http |
| 86.53.183.16:1080 | ✓ 1144ms | 否 | ✓ 1326ms | 否 | ✓ 1760ms | http |
| 14.225.212.37:7890 | ✓ 1275ms | ✓ 1652ms | 否 | ✓ 1687ms | 否 | http |
| 38.183.146.51:3125 | ✓ 1694ms | 否 | 否 | ✓ 1870ms | ✓ 1944ms | http |
| 24.144.86.173:1080 | ✓ 440ms | ✓ 1658ms | 否 | ✓ 1072ms | ✓ 1011ms | http |
| 45.149.92.147:5001 | ✓ 1736ms | 否 | ✓ 838ms | ✓ 1000ms | ✓ 815ms | http |
| 103.113.70.189:1081 | ✓ 1259ms | 否 | 否 | ✓ 1069ms | ✓ 964ms | http |
| 85.198.96.242:3128 | ✓ 586ms | 否 | 否 | ✓ 1577ms | ✓ 1216ms | http |
| 194.5.212.40:8080 | ✓ 418ms | ✓ 1595ms | ✓ 685ms | ✓ 1417ms | ✓ 1472ms | http |
| 168.235.110.63:3128 | ✓ 316ms | ✓ 1884ms | ✓ 1275ms | ✓ 1226ms | 否 | http |
| 88.80.150.82:8080 | ✓ 997ms | ✓ 1733ms | 否 | 否 | ✓ 1614ms | https |
| 34.101.184.164:3128 | ✓ 1656ms | 否 | ✓ 1094ms | ✓ 1497ms | ✓ 1153ms | http |
| 107.178.115.140:3128 | ✓ 737ms | ✓ 1657ms | ✓ 881ms | ✓ 1009ms | ✓ 812ms | http |
| 107.174.55.123:7878 | ✓ 728ms | 否 | 否 | ✓ 1053ms | ✓ 963ms | http |
| 144.24.29.128:10900 | ✓ 1492ms | ✓ 1969ms | ✓ 1989ms | ✓ 1372ms | ✓ 1566ms | http |
| 193.122.96.242:3128 | ✓ 1529ms | 否 | 否 | ✓ 1150ms | ✓ 921ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1928ms | ✓ 1163ms | ✓ 1480ms | 否 | http |
| 62.60.177.204:34094 | 否 | ✓ 1808ms | ✓ 1279ms | ✓ 1310ms | ✓ 901ms | http |
| 205.209.118.30:3138 | 否 | 否 | ✓ 1116ms | ✓ 1346ms | ✓ 944ms | http |
| 210.223.44.230:3128 | ✓ 1717ms | ✓ 1078ms | ✓ 834ms | 否 | ✓ 874ms | http |
| 150.249.255.91:3128 | ✓ 1467ms | 否 | ✓ 838ms | ✓ 1663ms | 否 | http |
| 165.227.5.10:8888 | 否 | 否 | ✓ 781ms | ✓ 1978ms | ✓ 747ms | http |
| 45.136.198.40:3128 | ✓ 1038ms | 否 | ✓ 1056ms | ✓ 1673ms | ✓ 1580ms | http |
| 120.92.212.16:8890 | ✓ 1127ms | ✓ 1506ms | ✓ 1301ms | ✓ 1814ms | ✓ 1168ms | http |
| 103.39.51.190:8080 | 否 | 否 | ✓ 1771ms | ✓ 1648ms | ✓ 1587ms | http |
| 38.145.208.94:8443 | ✓ 700ms | ✓ 897ms | ✓ 487ms | ✓ 1111ms | ✓ 736ms | http |
| 38.145.208.99:8443 | ✓ 703ms | ✓ 1310ms | ✓ 312ms | ✓ 978ms | ✓ 1499ms | http |
| 45.136.130.223:8443 | ✓ 714ms | ✓ 981ms | ✓ 891ms | ✓ 986ms | ✓ 1932ms | http |
| 121.230.9.26:1080 | ✓ 1691ms | ✓ 1938ms | 否 | 否 | ✓ 1486ms | http |
| 121.230.8.250:1080 | 否 | ✓ 1850ms | ✓ 1270ms | 否 | ✓ 1581ms | http |
| 61.52.131.172:8443 | ✓ 1092ms | ✓ 1383ms | ✓ 1086ms | ✓ 1355ms | ✓ 1087ms | http |
| 43.167.227.161:1080 | ✓ 1498ms | 否 | ✓ 1837ms | 否 | ✓ 1405ms | http |

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
