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

最后更新：2026-04-26 22:32:00 UTC（2026-04-27 06:32:00 UTC+8）

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
| 218.108.131.186:17890 | ✓ 1216ms | ✓ 1276ms | ✓ 1174ms | ✓ 1295ms | ✓ 1179ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1852ms | ✓ 1361ms | ✓ 1820ms | 否 | http |
| 47.84.59.16:1080 | ✓ 1213ms | ✓ 1909ms | ✓ 1094ms | ✓ 1307ms | ✓ 1048ms | http |
| 217.76.245.80:999 | ✓ 967ms | ✓ 1625ms | ✓ 1243ms | ✓ 1851ms | 否 | http |
| 64.188.67.154:1080 | ✓ 421ms | 否 | ✓ 918ms | 否 | ✓ 1434ms | http |
| 177.93.132.244:3128 | ✓ 1356ms | 否 | ✓ 1862ms | 否 | ✓ 1655ms | http |
| 62.113.119.14:8080 | ✓ 543ms | ✓ 1689ms | ✓ 898ms | ✓ 1427ms | ✓ 1067ms | http |
| 120.92.108.86:7890 | ✓ 1601ms | 否 | ✓ 1889ms | 否 | ✓ 1819ms | http |
| 103.157.200.126:3128 | ✓ 1674ms | 否 | ✓ 1723ms | 否 | ✓ 1668ms | http |
| 217.77.102.18:3128 | ✓ 1271ms | 否 | ✓ 1709ms | 否 | ✓ 1632ms | http |
| 47.85.51.197:1080 | ✓ 308ms | 否 | ✓ 1791ms | ✓ 1288ms | 否 | http |
| 211.95.152.50:45046 | 否 | ✓ 1733ms | ✓ 1360ms | ✓ 1525ms | ✓ 1241ms | http |
| 86.104.72.220:1081 | ✓ 276ms | ✓ 967ms | ✓ 975ms | ✓ 1453ms | ✓ 809ms | http |
| 77.110.119.136:3128 | 否 | 否 | ✓ 611ms | ✓ 1552ms | ✓ 959ms | http |
| 86.104.72.219:1081 | ✓ 319ms | ✓ 961ms | ✓ 613ms | ✓ 962ms | ✓ 1730ms | http |
| 46.101.95.183:8888 | ✓ 638ms | ✓ 1773ms | ✓ 1302ms | 否 | ✓ 1701ms | http |
| 101.32.244.83:8080 | ✓ 1171ms | ✓ 1658ms | ✓ 1153ms | ✓ 1669ms | 否 | http |
| 121.43.196.210:8222 | ✓ 1100ms | ✓ 1264ms | ✓ 1076ms | ✓ 1353ms | ✓ 1145ms | http |
| 121.43.196.213:8222 | ✓ 1162ms | ✓ 1269ms | ✓ 1103ms | ✓ 1339ms | ✓ 1170ms | http |
| 121.230.9.33:1080 | ✓ 1618ms | ✓ 1870ms | ✓ 1628ms | 否 | ✓ 1624ms | http |
| 152.42.177.32:8888 | ✓ 1194ms | 否 | ✓ 1592ms | ✓ 1589ms | ✓ 1507ms | http |
| 183.232.248.73:7890 | 否 | ✓ 1956ms | 否 | ✓ 1964ms | ✓ 1810ms | http |
| 121.230.8.136:1080 | ✓ 1328ms | ✓ 1721ms | ✓ 1224ms | ✓ 1851ms | 否 | http |
| 45.168.238.193:8443 | ✓ 1270ms | ✓ 1279ms | ✓ 927ms | ✓ 1888ms | ✓ 1055ms | http |
| 209.145.54.41:3128 | 否 | ✓ 1708ms | ✓ 1394ms | ✓ 1935ms | ✓ 1443ms | http |
| 120.92.212.16:7890 | ✓ 1196ms | 否 | ✓ 1913ms | ✓ 1387ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1880ms | 否 | 否 | ✓ 1942ms | ✓ 1819ms | http |
| 8.211.166.184:8081 | ✓ 1199ms | ✓ 1306ms | ✓ 917ms | ✓ 1086ms | ✓ 848ms | http |
| 86.104.74.110:1081 | ✓ 1006ms | ✓ 1954ms | ✓ 1328ms | 否 | 否 | http |
| 60.249.94.208:3128 | 否 | ✓ 1681ms | ✓ 1030ms | ✓ 1431ms | 否 | http |
| 210.223.44.230:3128 | ✓ 1620ms | ✓ 1307ms | 否 | 否 | ✓ 1717ms | http |
| 114.55.226.123:10086 | ✓ 1290ms | ✓ 1684ms | ✓ 1350ms | 否 | ✓ 1339ms | http |
| 61.52.131.172:8443 | ✓ 1105ms | 否 | ✓ 1181ms | ✓ 1403ms | 否 | http |
| 47.84.73.61:1080 | ✓ 1319ms | ✓ 1980ms | ✓ 1210ms | ✓ 1249ms | ✓ 1059ms | http |

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
