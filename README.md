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

最后更新：2026-03-04 18:40:16 UTC（2026-03-05 02:40:16 UTC+8）

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
| 205.209.118.30:3138 | ✓ 1204ms | ✓ 1088ms | ✓ 1124ms | ✓ 1202ms | ✓ 937ms | http |
| 125.128.12.144:3128 | ✓ 953ms | ✓ 1767ms | ✓ 1134ms | ✓ 1366ms | ✓ 1541ms | http |
| 14.56.177.44:3128 | ✓ 688ms | ✓ 954ms | ✓ 1227ms | ✓ 1068ms | ✓ 840ms | http |
| 61.72.110.54:3128 | ✓ 725ms | 否 | ✓ 1029ms | ✓ 1126ms | ✓ 957ms | http |
| 125.128.12.14:3128 | ✓ 1677ms | ✓ 1165ms | ✓ 974ms | ✓ 1108ms | ✓ 925ms | http |
| 162.240.154.26:3128 | ✓ 785ms | 否 | ✓ 1391ms | ✓ 1733ms | ✓ 1237ms | http |
| 61.72.221.194:3128 | ✓ 1728ms | 否 | ✓ 1243ms | 否 | ✓ 887ms | http |
| 14.56.107.244:3128 | ✓ 1718ms | 否 | ✓ 1342ms | 否 | ✓ 893ms | http |
| 121.128.121.54:3128 | ✓ 1689ms | ✓ 1455ms | ✓ 1994ms | 否 | ✓ 1450ms | http |
| 61.72.221.234:3128 | 否 | 否 | ✓ 1263ms | ✓ 1440ms | ✓ 1338ms | http |
| 192.166.82.55:1080 | ✓ 911ms | 否 | ✓ 1783ms | ✓ 1884ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1321ms | 否 | ✓ 999ms | ✓ 1526ms | ✓ 1018ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 1118ms | ✓ 1041ms | ✓ 844ms | http |
| 165.227.5.10:8888 | 否 | 否 | ✓ 698ms | ✓ 1322ms | ✓ 1061ms | http |
| 101.43.255.96:80 | ✓ 1175ms | ✓ 1267ms | ✓ 1022ms | ✓ 1285ms | ✓ 1036ms | http |
| 61.72.110.94:3128 | ✓ 757ms | 否 | ✓ 1452ms | ✓ 1613ms | ✓ 1403ms | http |
| 120.92.212.16:7890 | ✓ 1046ms | ✓ 1271ms | ✓ 1350ms | 否 | ✓ 1734ms | http |
| 61.72.221.94:3128 | ✓ 892ms | 否 | ✓ 1352ms | 否 | ✓ 929ms | http |
| 91.193.240.157:9877 | ✓ 1173ms | 否 | ✓ 1801ms | 否 | ✓ 1704ms | http |
| 144.31.25.69:21064 | ✓ 1148ms | 否 | ✓ 1273ms | 否 | ✓ 1951ms | http |
| 210.223.44.230:3128 | ✓ 888ms | ✓ 1054ms | ✓ 678ms | ✓ 1029ms | ✓ 791ms | http |
| 49.151.187.10:8082 | ✓ 1552ms | 否 | ✓ 1356ms | ✓ 1443ms | ✓ 1450ms | http |
| 35.234.17.221:8080 | ✓ 955ms | ✓ 1708ms | ✓ 1370ms | 否 | 否 | http |
| 207.254.71.62:8088 | ✓ 1136ms | ✓ 1810ms | ✓ 1404ms | 否 | ✓ 1617ms | http |
| 138.124.53.25:7443 | ✓ 980ms | ✓ 1683ms | ✓ 1539ms | ✓ 1679ms | ✓ 1645ms | http |
| 120.55.163.237:10086 | ✓ 954ms | ✓ 1120ms | ✓ 959ms | ✓ 1174ms | ✓ 943ms | http |
| 1.12.62.237:8080 | ✓ 1672ms | ✓ 1554ms | 否 | ✓ 1895ms | 否 | http |
| 81.70.169.194:80 | 否 | ✓ 1390ms | ✓ 1051ms | ✓ 1400ms | ✓ 1111ms | http |
| 45.140.147.155:1082 | ✓ 1249ms | 否 | 否 | ✓ 1699ms | ✓ 1120ms | http |
| 45.140.147.155:1081 | ✓ 706ms | 否 | ✓ 721ms | ✓ 1661ms | 否 | http |
| 103.35.188.243:3128 | 否 | ✓ 1080ms | 否 | ✓ 1229ms | ✓ 940ms | http |
| 62.113.119.14:8080 | ✓ 664ms | 否 | ✓ 1609ms | 否 | ✓ 1122ms | http |
| 103.215.36.88:17565 | ✓ 1180ms | ✓ 1410ms | ✓ 1199ms | ✓ 1499ms | ✓ 1138ms | http |
| 45.136.198.40:3128 | ✓ 1120ms | 否 | ✓ 1351ms | 否 | ✓ 1604ms | http |
| 159.89.31.62:8080 | ✓ 518ms | 否 | ✓ 1820ms | 否 | ✓ 1600ms | http |
| 109.234.38.35:3128 | ✓ 993ms | ✓ 1902ms | 否 | 否 | ✓ 1650ms | http |
| 46.249.103.192:443 | ✓ 701ms | 否 | ✓ 1299ms | ✓ 1514ms | 否 | http |
| 101.47.73.135:3128 | ✓ 1989ms | 否 | ✓ 1882ms | 否 | ✓ 1471ms | http |
| 103.215.36.88:13763 | ✓ 1158ms | ✓ 1492ms | ✓ 1304ms | ✓ 1422ms | ✓ 1126ms | http |
| 103.39.51.190:8080 | ✓ 1744ms | 否 | ✓ 1904ms | ✓ 1456ms | ✓ 1507ms | http |
| 74.48.78.224:2080 | ✓ 756ms | 否 | ✓ 1258ms | ✓ 964ms | ✓ 830ms | http |
| 211.171.114.154:3128 | 否 | ✓ 1403ms | ✓ 1794ms | ✓ 1573ms | ✓ 1020ms | http |
| 103.82.23.118:5188 | 否 | 否 | ✓ 1375ms | ✓ 1888ms | ✓ 1546ms | http |
| 172.212.68.37:3128 | ✓ 524ms | ✓ 1691ms | ✓ 878ms | ✓ 1473ms | ✓ 911ms | http |
| 103.67.46.225:3125 | 否 | 否 | ✓ 1879ms | ✓ 1682ms | ✓ 1707ms | http |
| 222.228.171.92:8080 | ✓ 832ms | 否 | ✓ 1844ms | ✓ 1972ms | ✓ 1860ms | http |

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
