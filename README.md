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

最后更新：2026-03-09 08:43:01 UTC（2026-03-09 16:43:01 UTC+8）

**代理总数：45**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 44 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 45 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 168.235.110.63:3128 | ✓ 1325ms | 否 | ✓ 813ms | ✓ 995ms | ✓ 740ms | http |
| 61.72.110.54:3128 | ✓ 762ms | ✓ 1768ms | 否 | ✓ 1225ms | ✓ 985ms | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1102ms | ✓ 1592ms | ✓ 1366ms | http |
| 120.92.212.16:7890 | ✓ 1113ms | ✓ 1438ms | ✓ 1124ms | 否 | 否 | http |
| 1.231.81.166:3128 | ✓ 1301ms | 否 | ✓ 1523ms | ✓ 1498ms | ✓ 1356ms | http |
| 165.227.5.10:8888 | ✓ 725ms | ✓ 1988ms | 否 | ✓ 1214ms | 否 | http |
| 121.128.121.54:3128 | ✓ 914ms | 否 | ✓ 1563ms | ✓ 1377ms | 否 | http |
| 202.155.12.161:443 | ✓ 1835ms | 否 | ✓ 1649ms | 否 | ✓ 1279ms | http |
| 178.236.245.59:3128 | ✓ 759ms | 否 | ✓ 1184ms | ✓ 1956ms | ✓ 1424ms | http |
| 178.236.245.17:3128 | ✓ 758ms | 否 | ✓ 1185ms | ✓ 1999ms | ✓ 1397ms | http |
| 162.243.149.86:31028 | ✓ 1311ms | ✓ 914ms | ✓ 1177ms | ✓ 1001ms | ✓ 902ms | http |
| 194.213.18.200:443 | ✓ 1278ms | 否 | ✓ 1863ms | 否 | ✓ 1793ms | http |
| 39.104.201.40:7890 | ✓ 1001ms | ✓ 1176ms | ✓ 983ms | ✓ 1186ms | ✓ 1859ms | http |
| 194.87.43.49:8888 | ✓ 1034ms | 否 | ✓ 1554ms | ✓ 1988ms | ✓ 1732ms | http |
| 101.43.255.96:80 | ✓ 1478ms | ✓ 1492ms | ✓ 1607ms | ✓ 1689ms | ✓ 1236ms | http |
| 115.231.181.40:8128 | ✓ 1129ms | ✓ 1385ms | 否 | 否 | ✓ 1087ms | http |
| 185.115.74.185:8080 | ✓ 1009ms | ✓ 1845ms | ✓ 1278ms | 否 | 否 | http |
| 125.128.12.14:3128 | ✓ 1358ms | ✓ 1615ms | ✓ 1193ms | ✓ 1186ms | ✓ 970ms | http |
| 159.89.31.62:8080 | ✓ 841ms | ✓ 1683ms | ✓ 1954ms | ✓ 1628ms | ✓ 1358ms | http |
| 14.56.107.244:3128 | 否 | 否 | ✓ 784ms | ✓ 1304ms | ✓ 947ms | http |
| 185.241.5.57:3128 | ✓ 1683ms | 否 | ✓ 1388ms | 否 | ✓ 1611ms | http |
| 81.70.169.194:80 | ✓ 1755ms | ✓ 1647ms | 否 | ✓ 1459ms | ✓ 1145ms | http |
| 67.169.98.211:443 | ✓ 1138ms | 否 | ✓ 586ms | ✓ 1309ms | ✓ 1517ms | http |
| 162.248.165.72:1080 | ✓ 1095ms | 否 | ✓ 1725ms | ✓ 1648ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1460ms | 否 | ✓ 1987ms | 否 | ✓ 1150ms | http |
| 35.225.22.61:80 | ✓ 960ms | 否 | ✓ 942ms | ✓ 1297ms | ✓ 875ms | http |
| 61.72.221.194:3128 | 否 | 否 | ✓ 769ms | ✓ 1257ms | ✓ 1210ms | http |
| 107.172.125.217:3128 | ✓ 1414ms | 否 | ✓ 912ms | ✓ 928ms | ✓ 775ms | http |
| 152.42.213.210:8080 | 否 | 否 | ✓ 1794ms | ✓ 1334ms | ✓ 1104ms | http |
| 14.225.222.164:7890 | 否 | ✓ 1932ms | ✓ 1519ms | 否 | ✓ 1541ms | http |
| 190.9.109.207:999 | ✓ 1234ms | ✓ 1351ms | ✓ 1088ms | ✓ 1528ms | ✓ 1185ms | http |
| 190.9.109.198:999 | ✓ 1234ms | ✓ 1405ms | ✓ 1142ms | ✓ 1404ms | ✓ 1215ms | http |
| 27.254.99.183:8118 | ✓ 1077ms | 否 | ✓ 1434ms | ✓ 1526ms | 否 | http |
| 49.151.178.218:8082 | ✓ 1662ms | 否 | 否 | ✓ 1541ms | ✓ 1578ms | http |
| 88.80.150.82:8080 | ✓ 1343ms | ✓ 1909ms | 否 | ✓ 1912ms | ✓ 1663ms | https |
| 121.230.8.211:1080 | ✓ 1100ms | ✓ 1931ms | ✓ 1145ms | 否 | ✓ 1171ms | http |
| 62.113.119.14:8080 | ✓ 574ms | ✓ 1595ms | ✓ 1793ms | ✓ 1460ms | ✓ 1821ms | http |
| 58.220.95.11:11023 | ✓ 1094ms | ✓ 1385ms | ✓ 1141ms | ✓ 1390ms | ✓ 1064ms | http |
| 45.136.198.40:3128 | ✓ 1430ms | ✓ 1782ms | ✓ 974ms | ✓ 1659ms | ✓ 1512ms | http |
| 167.172.253.162:4857 | 否 | ✓ 1397ms | ✓ 918ms | ✓ 1412ms | 否 | http |
| 64.186.232.4:10808 | 否 | 否 | ✓ 1170ms | ✓ 1084ms | ✓ 1469ms | http |
| 203.150.113.193:8080 | ✓ 1807ms | 否 | ✓ 1779ms | ✓ 1942ms | ✓ 1850ms | http |
| 103.39.51.190:8080 | ✓ 1926ms | 否 | 否 | ✓ 1687ms | ✓ 1700ms | http |
| 45.140.147.155:1082 | ✓ 448ms | 否 | ✓ 893ms | ✓ 1483ms | 否 | http |
| 45.140.147.155:1081 | ✓ 867ms | ✓ 1488ms | ✓ 1313ms | ✓ 1695ms | ✓ 1331ms | http |

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
